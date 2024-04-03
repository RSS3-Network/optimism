FROM --platform=$BUILDPLATFORM rss3/rollup-data-availability/da-rpc:latest as rust

RUN ls /lib
RUN ls /gopkg/da-rpc

# Copy lib from da-rpc
COPY --from=rust /gopkg/da-rpc/lib/ /usr/local/lib

ARG OP_STACK_GO_BUILDER=us-docker.pkg.dev/oplabs-tools-artifacts/images/op-stack-go:latest
FROM $OP_STACK_GO_BUILDER as builder
# See "make golang-docker" and /ops/docker/op-stack-go

FROM alpine:3.18

COPY --from=builder /usr/local/bin/op-batcher /usr/local/bin/op-batcher

ENV LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH

CMD ["op-batcher"]
