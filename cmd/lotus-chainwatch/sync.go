package main

import (
	"bytes"
	"container/list"
	"context"
	"sync"

	actors2 "github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/address"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"
)

func runSyncer(ctx context.Context, api api.FullNode, st *storage) {
	notifs, err := api.ChainNotify(ctx)
	if err != nil {
		panic(err)
	}
	go func() {
		for notif := range notifs {
			for _, change := range notif {
				switch change.Type {
				case store.HCCurrent:
					fallthrough
				case store.HCApply:
					syncHead(ctx, api, st, change.Val)
				case store.HCRevert:
					log.Warnf("revert todo")
				}
			}
		}
	}()
}

type minerKey struct {
	addr      address.Address
	act       types.Actor
	stateroot cid.Cid
}

type minerInfo struct {
	state actors2.StorageMinerActorState
	info  actors2.MinerInfo
}

func syncHead(ctx context.Context, api api.FullNode, st *storage, ts *types.TipSet) {
	addresses := map[address.Address]address.Address{}
	actors := map[address.Address]map[types.Actor]cid.Cid{}
	var alk sync.Mutex

	log.Infof("Getting headers / actors")

	toSync := map[cid.Cid]*types.BlockHeader{}
	toVisit := list.New()

	for _, header := range ts.Blocks() {
		toVisit.PushBack(header)
	}

	for toVisit.Len() > 0 {
		bh := toVisit.Remove(toVisit.Back()).(*types.BlockHeader)

		if _, seen := toSync[bh.Cid()]; seen || st.hasBlock(bh.Cid()) {
			continue
		}

		toSync[bh.Cid()] = bh
		addresses[bh.Miner] = address.Undef

		if len(toSync)%500 == 10 {
			log.Infof("todo: (%d) %s", len(toSync), bh.Cid())
		}

		if len(bh.Parents) == 0 {
			continue
		}

		pts, err := api.ChainGetTipSet(ctx, types.NewTipSetKey(bh.Parents...))
		if err != nil {
			log.Error(err)
			continue
		}

		for _, header := range pts.Blocks() {
			toVisit.PushBack(header)
		}
	}

	log.Infof("Syncing %d blocks", len(toSync))

	log.Infof("Persisting actors")

	paDone := 0
	par(50, maparr(toSync), func(bh *types.BlockHeader) {
		paDone++
		if paDone%100 == 0 {
			log.Infof("pa: %d %d%%", paDone, (paDone*100)/len(toSync))
		}

		if len(bh.Parents) == 0 { // genesis case
			ts, err := types.NewTipSet([]*types.BlockHeader{bh})
			aadrs, err := api.StateListActors(ctx, ts)
			if err != nil {
				log.Error(err)
				return
			}

			par(50, aadrs, func(addr address.Address) {
				act, err := api.StateGetActor(ctx, addr, ts)
				if err != nil {
					log.Error(err)
					return
				}
				alk.Lock()
				_, ok := actors[addr]
				if !ok {
					actors[addr] = map[types.Actor]cid.Cid{}
				}
				actors[addr][*act] = bh.ParentStateRoot
				addresses[addr] = address.Undef
				alk.Unlock()
			})

			return
		}

		pts, err := api.ChainGetTipSet(ctx, types.NewTipSetKey(bh.Parents...))
		if err != nil {
			log.Error(err)
			return
		}

		changes, err := api.StateChangedActors(ctx, pts.ParentState(), bh.ParentStateRoot)
		if err != nil {
			log.Error(err)
			return
		}

		for a, act := range changes {
			addr, err := address.NewFromString(a)
			if err != nil {
				log.Error(err)
				return
			}

			alk.Lock()
			_, ok := actors[addr]
			if !ok {
				actors[addr] = map[types.Actor]cid.Cid{}
			}
			actors[addr][act] = bh.ParentStateRoot
			addresses[addr] = address.Undef
			alk.Unlock()
		}
	})

	if err := st.storeActors(actors); err != nil {
		log.Error(err)
		return
	}

	log.Infof("Persisting miners")

	miners := map[minerKey]*minerInfo{}

	for addr, m := range actors {
		for actor, c := range m {
			if actor.Code != actors2.StorageMinerCodeCid {
				continue
			}

			miners[minerKey{
				addr:      addr,
				act:       actor,
				stateroot: c,
			}] = &minerInfo{}
		}
	}

	par(50, kvmaparr(miners), func(it func() (minerKey, *minerInfo)) {
		k, info := it()

		astb, err := api.ChainReadObj(ctx, k.act.Head)
		if err != nil {
			log.Error(err)
			return
		}

		if err := info.state.UnmarshalCBOR(bytes.NewReader(astb)); err != nil {
			log.Error(err)
			return
		}

		ib, err := api.ChainReadObj(ctx, info.state.Info)
		if err != nil {
			log.Error(err)
			return
		}

		if err := info.info.UnmarshalCBOR(bytes.NewReader(ib)); err != nil {
			log.Error(err)
			return
		}
	})

	if err := st.storeMiners(miners); err != nil {
		log.Error(err)
		return
	}

	log.Infof("Persisting headers")
	if err := st.storeHeaders(toSync, true); err != nil {
		log.Error(err)
		return
	}

	log.Infof("Getting messages")

	msgs, incls := fetchMessages(ctx, api, toSync)

	if err := st.storeMessages(msgs); err != nil {
		log.Error(err)
		return
	}

	if err := st.storeMsgInclusions(incls); err != nil {
		log.Error(err)
		return
	}

	log.Infof("Getting parent receipts")

	receipts := fetchParentReceipts(ctx, api, toSync)

	if err := st.storeReceipts(receipts); err != nil {
		log.Error(err)
		return
	}

	log.Infof("Resolving addresses")

	for _, message := range msgs {
		addresses[message.To] = address.Undef
		addresses[message.From] = address.Undef
	}

	par(50, kmaparr(addresses), func(addr address.Address) {
		raddr, err := api.StateLookupID(ctx, addr, nil)
		if err != nil {
			log.Warn(err)
			return
		}
		alk.Lock()
		addresses[addr] = raddr
		alk.Unlock()
	})

	if err := st.storeAddressMap(addresses); err != nil {
		log.Error(err)
		return
	}

	log.Infof("Sync done")
}

func fetchMessages(ctx context.Context, api api.FullNode, toSync map[cid.Cid]*types.BlockHeader) (map[cid.Cid]*types.Message, map[cid.Cid][]cid.Cid) {
	var lk sync.Mutex
	messages := map[cid.Cid]*types.Message{}
	inclusions := map[cid.Cid][]cid.Cid{} // block -> msgs

	par(50, maparr(toSync), func(header *types.BlockHeader) {
		msgs, err := api.ChainGetBlockMessages(ctx, header.Cid())
		if err != nil {
			log.Error(err)
			return
		}

		vmm := make([]*types.Message, 0, len(msgs.Cids))
		for _, m := range msgs.BlsMessages {
			vmm = append(vmm, m)
		}

		for _, m := range msgs.SecpkMessages {
			vmm = append(vmm, &m.Message)
		}

		lk.Lock()
		for _, message := range vmm {
			messages[message.Cid()] = message
			inclusions[header.Cid()] = append(inclusions[header.Cid()], message.Cid())
		}
		lk.Unlock()
	})

	return messages, inclusions
}

type mrec struct {
	msg   cid.Cid
	state cid.Cid
	idx   int
}

func fetchParentReceipts(ctx context.Context, api api.FullNode, toSync map[cid.Cid]*types.BlockHeader) map[mrec]*types.MessageReceipt {
	var lk sync.Mutex
	out := map[mrec]*types.MessageReceipt{}

	par(50, maparr(toSync), func(header *types.BlockHeader) {
		recs, err := api.ChainGetParentReceipts(ctx, header.Cid())
		if err != nil {
			log.Error(err)
			return
		}
		msgs, err := api.ChainGetParentMessages(ctx, header.Cid())
		if err != nil {
			log.Error(err)
			return
		}

		lk.Lock()
		for i, r := range recs {
			out[mrec{
				msg:   msgs[i].Cid,
				state: header.ParentStateRoot,
				idx:   i,
			}] = r
		}
		lk.Unlock()
	})

	return out
}
