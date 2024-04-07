FROM --platform=$BUILDPLATFORM rss3/rollup-data-availability/da-rpc:latest as rust

RUN ls /lib
RUN ls /gopkg/da-rpc

FROM rss3/op-stack-go:latest as builder
# See "make golang-docker" and /ops/docker/op-stack-go

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    jq \
    bash \
    curl

COPY --from=builder /usr/local/bin/op-proposer /usr/local/bin/op-proposer

# Copy lib from da-rpc
COPY --from=rust /gopkg/da-rpc/lib/ /usr/local/lib

ENV LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH
CMD ["op-proposer"]
