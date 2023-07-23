#!/bin/bash -e
ARCH=$(arch)

FILENAME="protoc-21.5-linux-${ARCH}.zip"

if [ "${ARCH}" = 'aarch64' ]; then
    ARCH="aarch_64"
fi

curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v21.5/protoc-21.5-linux-${ARCH}.zip
unzip protoc-21.5-linux-${ARCH}.zip -d /home/vscode/.local
