variable "REGISTRY" {
  default = "docker.io"
}

variable "REPOSITORY" {
  default = "rss3"
}

variable "GIT_COMMIT" {
  default = "dev"
}

variable "GIT_DATE" {
  default = "0"
}

variable "GIT_VERSION" {
  default = "docker"  // original default as set in proxyd file, not used by full go stack, yet
}

variable "IMAGE_TAGS" {
  default = "${GIT_COMMIT}" // split by ","
}

variable "PLATFORMS" {
  // You can override this as "linux/amd64,linux/arm64".
  // Only a specify a single platform when `--load` ing into docker.
  // Multi-platform is supported when outputting to disk or pushing to a registry.
  // Multi-platform builds can be tested locally with:  --set="*.output=type=image,push=false"
  default = "linux/amd64"
}

group "validate" {
  targets = ["op-node", "op-batcher", "op-proposer"]
}

target "op-stack-go" {
  dockerfile = "ops/docker/op-stack-go/Dockerfile"
  context = "."
  args = {
    GIT_COMMIT = "${GIT_COMMIT}"
    GIT_DATE = "${GIT_DATE}"
  }
  platforms = split(",", PLATFORMS)
  tags = [for tag in split(",", IMAGE_TAGS) : "${REGISTRY}/${REPOSITORY}/op-stack-go:${tag}"]
}

target "op-node" {
  dockerfile = "Dockerfile"
  context = "./op-node"
  args = {
    OP_STACK_GO_BUILDER = "op-stack-go"
  }
#  contexts = {
#    op-stack-go: "target:op-stack-go"
#  }
  platforms = split(",", PLATFORMS)
  tags = [for tag in split(",", IMAGE_TAGS) : "${REGISTRY}/${REPOSITORY}/op-node:${tag}"]
}

target "op-batcher" {
  dockerfile = "Dockerfile"
  context = "./op-batcher"
  args = {
    OP_STACK_GO_BUILDER = "op-stack-go"
  }
#  contexts = {
#    op-stack-go: "target:op-stack-go"
#  }
  platforms = split(",", PLATFORMS)
  tags = [for tag in split(",", IMAGE_TAGS) : "${REGISTRY}/${REPOSITORY}/op-batcher:${tag}"]
}

target "op-proposer" {
  dockerfile = "Dockerfile"
  context = "./op-proposer"
  args = {
    OP_STACK_GO_BUILDER = "op-stack-go"
  }
#  contexts = {
#    op-stack-go: "target:op-stack-go"
#  }
  platforms = split(",", PLATFORMS)
  tags = [for tag in split(",", IMAGE_TAGS) : "${REGISTRY}/${REPOSITORY}/op-proposer:${tag}"]
}

target "op-heartbeat" {
  dockerfile = "Dockerfile"
  context = "./op-heartbeat"
  args = {
    OP_STACK_GO_BUILDER = "op-stack-go"
  }
  contexts = {
    op-stack-go: "target:op-stack-go"
  }
  platforms = split(",", PLATFORMS)
  tags = [for tag in split(",", IMAGE_TAGS) : "${REGISTRY}/${REPOSITORY}/op-heartbeat:${tag}"]
}

target "op-ufm" {
  dockerfile = "./op-ufm/Dockerfile"
  context = "./"
  args = {
    // op-ufm dockerfile has no _ in the args
    GITCOMMIT = "${GIT_COMMIT}"
    GITDATE = "${GIT_DATE}"
    GITVERSION = "${GIT_VERSION}"
  }
  platforms = split(",", PLATFORMS)
  tags = [for tag in split(",", IMAGE_TAGS) : "${REGISTRY}/${REPOSITORY}/op-ufm:${tag}"]
}