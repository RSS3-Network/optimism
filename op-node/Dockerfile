ARG OP_STACK_GO_BUILDER=rss3/op-stack-go:latest
FROM $OP_STACK_GO_BUILDER as builder
# See "make golang-docker" and /ops/docker/op-stack-go

FROM --platform=$BUILDPLATFORM rss3/da-rpc:latest as rust

RUN ls /lib
RUN ls /gopkg/da-rpc

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    jq \
    bash \
    curl

COPY --from=builder /usr/local/bin/op-node /usr/local/bin/op-node

# Copy lib from da-rpc
COPY --from=rust /gopkg/da-rpc/lib/ /usr/local/lib

ENV LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH

CMD ["op-node"]
