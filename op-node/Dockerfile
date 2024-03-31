FROM --platform=$BUILDPLATFORM rss3/rollup-data-availability/da-rpc:latest as rust

RUN ls /lib
RUN ls /gopkg/da-rpc

FROM --platform=$BUILDPLATFORM golang:1.21 as builder

ARG VERSION=v0.0.0

RUN apt-get update && apt-get install -y \
    git \
    jq \
    make \
    bash \
    openssl \
    libssl-dev \
    protobuf-compiler \
    pkg-config \
    cbindgen \
    clang \
    curl

# build op-node with the shared go.mod & go.sum files
COPY ./op-node /app/op-node
COPY ./op-chain-ops /app/op-chain-ops
COPY ./op-service /app/op-service
COPY ./op-conductor /app/op-conductor
COPY ./op-plasma  /app/op-plasma
COPY ./op-bindings /app/op-bindings
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum
COPY ./Makefile /app/Makefile
COPY ./.git /app/.git


# Copy da rpc
COPY --from=rust /gopkg/da-rpc/lib/ /usr/local/lib
COPY ./go.work /app/go.work

WORKDIR /app
COPY ./Makefile /app/Makefile
RUN ls /app/Makefile

WORKDIR /app/

RUN ls /app/

RUN go mod download

ARG TARGETOS TARGETARCH

RUN make op-node VERSION="$VERSION" GOOS=$TARGETOS GOARCH=$TARGETARCH

RUN mv /app/op-node/bin/op-node /

FROM debian:bookworm-slim

COPY --from=rust /gopkg/da-rpc/lib/ /usr/local/lib
COPY --from=builder /op-node /usr/local/bin

RUN apt-get update && apt-get install -y \
    openssl \
    libssl-dev \
    ca-certificates \
    protobuf-compiler

ENV LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH

CMD ["op-node"]
