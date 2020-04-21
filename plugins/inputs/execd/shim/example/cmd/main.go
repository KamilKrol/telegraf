package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	// TODO: import your plugins
	// _ "github.com/my_github_user/my_plugin_repo/plugins/inputs/my_plugin_name"

	"github.com/influxdata/telegraf/config"
	"github.com/influxdata/telegraf/plugins/inputs/execd/shim"
)

var pollInterval = 1 * time.Second // can use shim.PollIntervalDisabled

var configFile = flag.String("config", "", "path to the config file for this plugin")
var cfg *config.Config

func main() {
	flag.Parse()

	cfg = config.NewConfig()

	if *configFile != "" {
		if err := cfg.LoadConfig(*configFile); err != nil {
			fmt.Fprintf(os.Stderr, "Err: %s\n", err)
			os.Exit(1)
		}
	}

	if err := shim.RunPlugins(cfg, pollInterval); err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s\n", err)
		os.Exit(1)
	}
}
