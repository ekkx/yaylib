# Containerized development environment for the SDK and its bundled mock
# server. Mirrors the CI matrix (Ubuntu 24.04 / Go / Node 22 /
# Python 3.12) so a contributor can build, test and run everything
# without installing any toolchain on the host.
FROM ubuntu:24.04

# Go is copied from the official image so the version is pinned without
# guessing a download URL.
COPY --from=golang:1.25-bookworm /usr/local/go /usr/local/go

RUN apt-get update \
 && apt-get install -y --no-install-recommends \
      ca-certificates curl git bash build-essential \
      python3 python3-venv python3-pip \
 && curl -fsSL https://deb.nodesource.com/setup_22.x | bash - \
 && apt-get install -y --no-install-recommends nodejs \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

ENV PATH=/usr/local/go/bin:/go/bin:$PATH \
    GOPATH=/go \
    GOMODCACHE=/cache/go/mod \
    GOCACHE=/cache/go/build \
    NPM_CONFIG_CACHE=/cache/npm \
    PIP_CACHE_DIR=/cache/pip \
    PYTHONDONTWRITEBYTECODE=1

WORKDIR /workspace

# Interactive shell by default; the compose services override this with
# the concrete command to run.
CMD ["bash"]
