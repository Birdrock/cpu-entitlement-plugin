package plugins

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"code.cloudfoundry.org/cli/cf/terminal"
	"code.cloudfoundry.org/cli/cf/trace"
	"code.cloudfoundry.org/cli/plugin"
	"code.cloudfoundry.org/cpu-entitlement-plugin/cf"
	"code.cloudfoundry.org/cpu-entitlement-plugin/fetchers"
	"code.cloudfoundry.org/cpu-entitlement-plugin/output"
	"code.cloudfoundry.org/cpu-entitlement-plugin/reporter"
	"code.cloudfoundry.org/lager/v3"
	flags "github.com/jessevdk/go-flags"
)

type CPUEntitlementAdminPlugin struct{ Version string }

func NewOverEntitlementInstancesPlugin() CPUEntitlementAdminPlugin {
	return CPUEntitlementAdminPlugin{}
}

func (p CPUEntitlementAdminPlugin) WithVersion(version string) CPUEntitlementAdminPlugin {
	p.Version = version
	return p
}

func (p CPUEntitlementAdminPlugin) Run(cli plugin.CliConnection, args []string) {
	traceLogger := trace.NewLogger(os.Stdout, true, os.Getenv("CF_TRACE"), "")
	ui := terminal.NewUI(os.Stdin, os.Stdout, terminal.NewTeePrinter(os.Stdout), traceLogger)

	opts := struct {
		Debug bool `short:"d" long:"debug" description:"Show verbose debug information"`
	}{}

	args, err := flags.ParseArgs(&opts, args)
	if err != nil {
		ui.Failed("Invalid arguments.")
		os.Exit(1)
	}

	logger := lager.NewLogger("over-entitlement-instances")
	outputSink := ioutil.Discard
	if opts.Debug {
		outputSink = os.Stdout
	}

	logger.RegisterSink(lager.NewPrettySink(outputSink, lager.DEBUG))

	logger.Info("start")
	defer logger.Info("end")
	if args[0] == "CLI-MESSAGE-UNINSTALL" {
		os.Exit(0)
	}

	logCacheURL, err := getLogCacheURL(cli)
	if err != nil {
		ui.Failed(err.Error())
		os.Exit(1)
	}

	ui.Warn("Note: This plugin is experimental.")

	sslIsDisabled, err := cli.IsSSLDisabled()
	if err != nil {
		ui.Failed(err.Error())
		os.Exit(1)
	}

	fetcher := fetchers.NewCumulativeUsageFetcher(createLogClient(logCacheURL, cli.AccessToken, sslIsDisabled))
	cfClient := cf.NewClient(cli, fetchers.NewProcessInstanceIDFetcher(createLogClient(logCacheURL, cli.AccessToken, sslIsDisabled)))
	reporter := reporter.NewOverEntitlementInstances(cfClient, fetcher)
	renderer := output.NewOverEntitlementInstancesRenderer(output.NewTerminalDisplay(ui))
	runner := NewOverEntitlementInstancesRunner(reporter, renderer)

	err = runner.Run(logger)
	if err != nil {
		ui.Failed(err.Error())
		os.Exit(1)
	}
}

func (p CPUEntitlementAdminPlugin) GetMetadata() plugin.PluginMetadata {
	major := 0
	minor := 0
	build := 0
	if p.Version != "" {
		version := strings.Split(p.Version, ".")
		if len(version) == 3 {
			major, _ = strconv.Atoi(version[0])
			minor, _ = strconv.Atoi(version[1])
			build, _ = strconv.Atoi(version[2])
		}
	}
	return plugin.PluginMetadata{
		Name: "CPUEntitlementAdminPlugin",
		Version: plugin.VersionType{
			Major: major,
			Minor: minor,
			Build: build,
		},
		Commands: []plugin.Command{
			{
				Name:     "over-entitlement-instances",
				Alias:    "oei",
				HelpText: "See which instances are over entitlement",
				UsageDetails: plugin.Usage{
					Usage: "cf over-entitlement-instances",
				},
			},
		},
	}
}
