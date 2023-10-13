package main

import (
	"code.cloudfoundry.org/cli/plugin"
	"code.cloudfoundry.org/cpu-entitlement-plugin/plugins"
)

var Version string

func main() {
	plugin.Start(plugins.NewCPUEntitlementPlugin().WithVersion(Version))
}
