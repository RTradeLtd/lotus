# Set up Golang build environment
FROM golang:1.13 AS build-env

# Install build dependencies
RUN apt-get update && apt-get install -y sudo curl git mesa-opencl-icd ocl-icd-opencl-dev
# install golang
RUN sudo apt install -y gcc git bzr jq pkg-config mesa-opencl-icd ocl-icd-opencl-dev
# download repo
RUN git clone https://github.com/filecoin-project/lotus.git
RUN cd lotus && make clean all && sudo make install
COPY lotus_docker_config.toml /root/.lotus/config.toml
EXPOSE 1234/tcp
EXPOSE 1235/tcp
ENTRYPOINT [ "lotus", "daemon" ]