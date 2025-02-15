# Setup Troubleshooting

Here is a command that will delete your chain data and any miners you have set up:

```sh
rm -rf ~/.lotus ~/.lotusstorage
```

This command usually resolves any issues with running `lotus` but it is not always required for updates. We will share information about when resetting your chain data and miners is required for an update in the future.

## Lotus daemon problems

```sh
WARN  peermgr peermgr/peermgr.go:131  failed to connect to bootstrap peer: failed to dial : all dials failed
  * [/ip4/147.75.80.17/tcp/1347] failed to negotiate security protocol: connected to wrong peer
```

* Try running the build steps again and make sure that you have the latest code from GitHub.

```sh
ERROR hello hello/hello.go:81 other peer has different genesis!
```

* Try deleting your file system's `~/.lotus` directory. Check that it exists with `ls ~/.lotus`.

```sh
- repo is already locked
```

* You already have another lotus deamon running.

## Failed messages

Some errors will occur that do not prevent Lotus from working:

```sh
ERROR chainstore  store/store.go:564  get message get failed: <Data CID>: blockstore: block not found

```

* Someone is requesting a **Data CID** from you that you don't have.