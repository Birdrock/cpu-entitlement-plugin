#!/bin/bash

set -eu
set -o pipefail

function run() {
    local arch os output
    arch="${1:?Please provide an architecture}"
    os="${2:?Please provide an OS}"
    version="${3:?Please provide a version}"
    output="${4:?Please provide an output directory}"

    GOARCH="${arch}" GOOS="${os}" \
        go build \
        -o "${output}/cpu-entitlement-plugin-${os}-${arch}" \
        -ldflags "-X main.Version=${version}" \
        ./cmd/cpu-entitlement

    GOARCH="${arch}" GOOS="${os}" \
        go build \
        -o "${output}/cpu-overentitlement-instances-plugin-${os}-${arch}" \
        -ldflags "-X main.Version=${version}" \
        ./cmd/cpu-overentitlement-instances
}

run "$@"
