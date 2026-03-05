package logcheck

import (
	"github.com/Razzle131/loglint/config"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("loglint", NewPlugin)
}

type Plugin struct {
}

func NewPlugin(settings any) (register.LinterPlugin, error) {
	// The configuration type will be map[string]any or []interface, it depends on your configuration.
	// You can use https://github.com/go-viper/mapstructure to convert map to struct.

	cfgPath, err := register.DecodeSettings[string](settings)
	if err != nil {
		return nil, err
	}

	cfg := config.Load(cfgPath)
	_ = New(cfg)

	return &Plugin{}, nil
}

func (f *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		{
			Name: "todo",
			Doc:  "find todos without an author",
			Run:  f.run,
		},
	}, nil
}

func (f *Plugin) GetLoadMode() string {
	// NOTE: the mode can be `register.LoadModeSyntax` or `register.LoadModeTypesInfo`.
	// - `register.LoadModeSyntax`: if the linter doesn't use types information.
	// - `register.LoadModeTypesInfo`: if the linter uses types information.

	return register.LoadModeTypesInfo
}

func (f *Plugin) run(pass *analysis.Pass) (any, error) {
	return run(pass)
}
