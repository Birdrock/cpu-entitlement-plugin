#!/bin/bash

set -eu
set -o pipefail

configure_cpu_entitlement_plugin
configure_cpu_overentitlement_instances_plugin

go run github.com/onsi/ginkgo/v2/ginkgo ${@}
