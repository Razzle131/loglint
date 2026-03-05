package plugin

import (
	"github.com/Razzle131/loglint/config"
	"github.com/Razzle131/loglint/logcheck"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("loglint", New)
}

type Plugin struct {
}

type PluginSettings struct {
	CfgPath string `json:"cfgPath"`
}

func New(settings any) (register.LinterPlugin, error) {
	s, err := register.DecodeSettings[PluginSettings](settings)
	if err != nil {
		return nil, err
	}

	cfg := config.Load(s.CfgPath)
	_ = logcheck.NewAnalyzer(cfg)

	return &Plugin{}, nil
}

func (f *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		{
			Name: "loglint",
			Doc:  "checks logging calls",
			Run:  f.run,
		},
	}, nil
}

func (f *Plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}

func (f *Plugin) run(pass *analysis.Pass) (any, error) {
	return logcheck.Run(pass)
}
