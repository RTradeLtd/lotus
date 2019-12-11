# stage 1 build - compiles the lotus node
FROM golang:1.13 AS build-env
RUN apt-get update && apt-get install -y sudo curl git mesa-opencl-icd ocl-icd-opencl-dev
RUN sudo apt install -y gcc git bzr jq pkg-config mesa-opencl-icd ocl-icd-opencl-dev
RUN git clone https://github.com/filecoin-project/lotus.git
RUN cd lotus && make clean all && sudo make install
# stage 2 build - installs dependencies, copies lotus binary from stage 1, and adds default config
FROM golang:1.13
RUN apt-get update -y && apt-get install -y mesa-opencl-icd ocl-icd-opencl-dev
COPY --from=build-env /usr/local/bin/lotus /usr/local/bin/lotus
COPY lotus_docker_config.toml /root/.lotus/config.toml
EXPOSE 1234/tcp
EXPOSE 1235/tcp
ENTRYPOINT [ "lotus", "daemon" ]