package blocksync

import (
	"fmt"
	"io"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

/* This file was generated by github.com/whyrusleeping/cbor-gen */

var _ = xerrors.Errorf

func (t *BlockSyncRequest) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.t.Start ([]cid.Cid) (slice)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Start)))); err != nil {
		return err
	}
	for _, v := range t.Start {
		if err := cbg.WriteCid(w, v); err != nil {
			return xerrors.Errorf("failed writing cid field t.Start: %w", err)
		}
	}

	// t.t.RequestLength (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.RequestLength))); err != nil {
		return err
	}

	// t.t.Options (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Options))); err != nil {
		return err
	}
	return nil
}

func (t *BlockSyncRequest) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.t.Start ([]cid.Cid) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Start: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.Start = make([]cid.Cid, extra)
	}
	for i := 0; i < int(extra); i++ {

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("reading cid field t.Start failed: %w", err)
		}
		t.Start[i] = c
	}

	// t.t.RequestLength (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.RequestLength = uint64(extra)
	// t.t.Options (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Options = uint64(extra)
	return nil
}

func (t *BlockSyncResponse) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.t.Chain ([]*blocksync.BSTipSet) (slice)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Chain)))); err != nil {
		return err
	}
	for _, v := range t.Chain {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.t.Status (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Status))); err != nil {
		return err
	}

	// t.t.Message (string) (string)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.Message)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.Message)); err != nil {
		return err
	}
	return nil
}

func (t *BlockSyncResponse) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.t.Chain ([]*blocksync.BSTipSet) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Chain: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.Chain = make([]*BSTipSet, extra)
	}
	for i := 0; i < int(extra); i++ {

		var v BSTipSet
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Chain[i] = &v
	}

	// t.t.Status (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.Status = uint64(extra)
	// t.t.Message (string) (string)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		t.Message = string(sval)
	}
	return nil
}

func (t *BSTipSet) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{133}); err != nil {
		return err
	}

	// t.t.Blocks ([]*types.BlockHeader) (slice)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.Blocks)))); err != nil {
		return err
	}
	for _, v := range t.Blocks {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.t.BlsMessages ([]*types.Message) (slice)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.BlsMessages)))); err != nil {
		return err
	}
	for _, v := range t.BlsMessages {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.t.BlsMsgIncludes ([][]uint64) (slice)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.BlsMsgIncludes)))); err != nil {
		return err
	}
	for _, v := range t.BlsMsgIncludes {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(v)))); err != nil {
			return err
		}
		for _, v := range v {
			if err := cbg.CborWriteHeader(w, cbg.MajUnsignedInt, v); err != nil {
				return err
			}
		}
	}

	// t.t.SecpkMessages ([]*types.SignedMessage) (slice)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.SecpkMessages)))); err != nil {
		return err
	}
	for _, v := range t.SecpkMessages {
		if err := v.MarshalCBOR(w); err != nil {
			return err
		}
	}

	// t.t.SecpkMsgIncludes ([][]uint64) (slice)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(t.SecpkMsgIncludes)))); err != nil {
		return err
	}
	for _, v := range t.SecpkMsgIncludes {
		if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajArray, uint64(len(v)))); err != nil {
			return err
		}
		for _, v := range v {
			if err := cbg.CborWriteHeader(w, cbg.MajUnsignedInt, v); err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *BSTipSet) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 5 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.t.Blocks ([]*types.BlockHeader) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Blocks: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.Blocks = make([]*types.BlockHeader, extra)
	}
	for i := 0; i < int(extra); i++ {

		var v types.BlockHeader
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.Blocks[i] = &v
	}

	// t.t.BlsMessages ([]*types.Message) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.BlsMessages: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.BlsMessages = make([]*types.Message, extra)
	}
	for i := 0; i < int(extra); i++ {

		var v types.Message
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.BlsMessages[i] = &v
	}

	// t.t.BlsMsgIncludes ([][]uint64) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.BlsMsgIncludes: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.BlsMsgIncludes = make([][]uint64, extra)
	}
	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}

			if extra > cbg.MaxLength {
				return fmt.Errorf("t.BlsMsgIncludes[i]: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}
			if extra > 0 {
				t.BlsMsgIncludes[i] = make([]uint64, extra)
			}
			for j := 0; j < int(extra); j++ {

				maj, val, err := cbg.CborReadHeader(br)
				if err != nil {
					return xerrors.Errorf("failed to read uint64 for t.BlsMsgIncludes[i] slice: %w", err)
				}

				if maj != cbg.MajUnsignedInt {
					return xerrors.Errorf("value read for array t.BlsMsgIncludes[i] was not a uint, instead got %d", maj)
				}

				t.BlsMsgIncludes[i][j] = val
			}

		}
	}

	// t.t.SecpkMessages ([]*types.SignedMessage) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.SecpkMessages: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.SecpkMessages = make([]*types.SignedMessage, extra)
	}
	for i := 0; i < int(extra); i++ {

		var v types.SignedMessage
		if err := v.UnmarshalCBOR(br); err != nil {
			return err
		}

		t.SecpkMessages[i] = &v
	}

	// t.t.SecpkMsgIncludes ([][]uint64) (slice)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.SecpkMsgIncludes: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}
	if extra > 0 {
		t.SecpkMsgIncludes = make([][]uint64, extra)
	}
	for i := 0; i < int(extra); i++ {
		{
			var maj byte
			var extra uint64
			var err error

			maj, extra, err = cbg.CborReadHeader(br)
			if err != nil {
				return err
			}

			if extra > cbg.MaxLength {
				return fmt.Errorf("t.SecpkMsgIncludes[i]: array too large (%d)", extra)
			}

			if maj != cbg.MajArray {
				return fmt.Errorf("expected cbor array")
			}
			if extra > 0 {
				t.SecpkMsgIncludes[i] = make([]uint64, extra)
			}
			for j := 0; j < int(extra); j++ {

				maj, val, err := cbg.CborReadHeader(br)
				if err != nil {
					return xerrors.Errorf("failed to read uint64 for t.SecpkMsgIncludes[i] slice: %w", err)
				}

				if maj != cbg.MajUnsignedInt {
					return xerrors.Errorf("value read for array t.SecpkMsgIncludes[i] was not a uint, instead got %d", maj)
				}

				t.SecpkMsgIncludes[i][j] = val
			}

		}
	}

	return nil
}
