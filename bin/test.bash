#!/bin/bash

set -eu
set -o pipefail

go run github.com/onsi/ginkgo/v2/ginkgo --skip-package e2e,integration ${@}
