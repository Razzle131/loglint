package main

import (
	"flag"

	"github.com/Razzle131/loglint/config"
	"github.com/Razzle131/loglint/logcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "server configuration file")
	flag.Parse()
	cfg := config.MustLoad(configPath)

	analyzer := logcheck.New(cfg)
	singlechecker.Main(analyzer)
}
