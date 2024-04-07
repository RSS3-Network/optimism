FROM iavlx/op-stack-go:latest as builder
# See "make golang-docker" and /ops/docker/op-stack-go

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y \
    jq \
    bash \
    curl

COPY --from=builder /usr/local/bin/op-proposer /usr/local/bin/op-proposer

CMD ["op-proposer"]
