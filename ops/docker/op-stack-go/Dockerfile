# FROM --platform=$BUILDPLATFORM rss3/rollup-data-availability/da-rpc:latest as rust
FROM --platform=$BUILDPLATFORM rss3/da-rpc:latest as rust

RUN ls /lib
RUN ls /gopkg/da-rpc

FROM --platform=$BUILDPLATFORM golang:1.22.7 as builder

RUN apt-get update && apt-get install -y \
    curl \
    tar \
    gzip \
    make \
    gcc \
    openssl \
    libssl-dev \
    musl-dev \
    git \
    jq \
    bash

RUN curl https://mise.run | MISE_INSTALL_PATH=/usr/local/bin/mise sh

# Copy lib from da-rpc
COPY --from=rust /gopkg/da-rpc/lib/ /usr/local/lib
ENV LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH

ARG TARGETARCH

# Install yq
RUN wget https://github.com/mikefarah/yq/releases/latest/download/yq_linux_$TARGETARCH -O /usr/local/bin/yq && \
    chmod +x /usr/local/bin/yq

# Install versioned toolchain
COPY ./mise.toml .
RUN mise trust && mise install -v -y just && cp $(mise which just) /usr/local/bin/just && just --version

# We copy the go.mod/sum first, so the `go mod download` does not have to re-run if dependencies do not change.
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

WORKDIR /app

RUN echo "go mod cache: $(go env GOMODCACHE)"
RUN echo "go build cache: $(go env GOCACHE)"

# warm-up the cache
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build go mod download

# NOTE: the Dockerfile.dockerignore file effectively describes all dependencies
COPY . /app

# We avoid copying the full .git dir into the build for just some metadata.
# Instead, specify:
# --build-arg GIT_COMMIT=$(git rev-parse HEAD)
# --build-arg GIT_DATE=$(git show -s --format='%ct')
ARG GIT_COMMIT
ARG GIT_DATE

ARG TARGETOS

# separate docker-builds:
# - proxyd
# - any JS/TS/smart-contract builds

# Build the Go services, utilizing caches and share the many common packages.
# The "id" defaults to the value of "target", the cache will thus be reused during this build.
# "sharing" defaults to "shared", the cache will thus be available to other concurrent docker builds.
ARG OP_NODE_VERSION=v0.0.0
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build cd op-node && make op-node  \
  GOOS=$TARGETOS GOARCH=$TARGETARCH GITCOMMIT=$GIT_COMMIT GITDATE=$GIT_DATE VERSION="$OP_NODE_VERSION"

ARG OP_BATCHER_VERSION=v0.0.0
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build cd op-batcher && make op-batcher  \
  GOOS=$TARGETOS GOARCH=$TARGETARCH GITCOMMIT=$GIT_COMMIT GITDATE=$GIT_DATE  VERSION="$OP_BATCHER_VERSION"

ARG OP_PROPOSER_VERSION=v0.0.0
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build cd op-proposer && make op-proposer  \
  GOOS=$TARGETOS GOARCH=$TARGETARCH GITCOMMIT=$GIT_COMMIT GITDATE=$GIT_DATE  VERSION="$OP_PROPOSER_VERSION"

FROM debian:bookworm-slim

COPY --from=builder /app/op-node/bin/op-node /usr/local/bin/
COPY --from=builder /app/op-batcher/bin/op-batcher /usr/local/bin/
COPY --from=builder /app/op-proposer/bin/op-proposer /usr/local/bin/
