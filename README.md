![Lotus](documentation/images/lotus_logo_h.png)

# Project Lotus - 莲

Lotus is an experimental implementation of the Filecoin Distributed Storage Network. For more details about Filecoin, check out the [Filecoin Spec](https://github.com/filecoin-project/specs).

## Development

All work is tracked via issues. An attempt at keeping an up-to-date view on remaining work is in the [lotus testnet github project board](https://github.com/filecoin-project/lotus/projects/1).

## Building & Documentation

For instructions on how to build lotus from source, please visit [https://docs.lotu.sh](https://docs.lotu.sh) or read the source [here](https://github.com/filecoin-project/lotus/tree/master/documentation).

## Docker

The docker image is intended to run the lotus node exposing the API for consumption by applications. The docker image also includes an NGINX reverse proxy to handle CORS processing. To run the docker image the suggested command is:

```shell
$> docker run -p "8080:8080" -p "1235:1235" -d rtradetech/lotus:latest
```

Where `8080` is mapping the NGINX reverse proxy port for the API, and `1235` is mapping the lotus node "swarm port".

## License

Dual-licensed under [MIT](https://github.com/filecoin-project/lotus/blob/master/LICENSE-MIT) + [Apache 2.0](https://github.com/filecoin-project/lotus/blob/master/LICENSE-APACHE)
