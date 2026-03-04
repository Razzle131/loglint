package main

import (
	"flag"

	"github.com/Razzle131/loglint/config"
	"github.com/Razzle131/loglint/logcheck"
	_ "go.uber.org/zap"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.yaml", "linter configuration file")
	flag.Parse()
	cfg := config.Load(configPath)

	analyzer := logcheck.New(cfg)
	singlechecker.Main(analyzer)
}
