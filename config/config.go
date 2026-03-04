package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

// Func describes a function to analyze, e.g. [slog.Info].
type Func struct {
	Name   string `yaml:"name"`   // full function name (example: (*log/slog.Logger).Info)
	MsgPos int    `yaml:"msgPos"` // position of msg argument starting from 0
	ArgPos int    `yaml:"argPos"` // The position of the "args ...any" argument starting from 0
}

type Config struct {
	EnabledRules int      `yaml:"enabled_rules" env:"ENABLED_RULES"` // bit mask for rules on/off (example: 1101=13 rule 2 is turned off)
	AvoidedData  []string `yaml:"avoided_data" env:"AVOIDED_DATA"`   // slice of avoided names, lowercase only
	EnabledFuncs []Func   `yaml:"enabled_funcs"`                     // slice of func definitions that linter searches for
}

const allRulesEnabledMask = 15

var defaultConfig = Config{
	EnabledRules: allRulesEnabledMask,
	AvoidedData:  []string{"password", "apikey", "token"},
	EnabledFuncs: []Func{
		{Name: "log/slog.Debug", MsgPos: 0, ArgPos: 1},
		{Name: "log/slog.Info", MsgPos: 0, ArgPos: 1},
		{Name: "log/slog.Warn", MsgPos: 0, ArgPos: 1},
		{Name: "log/slog.Error", MsgPos: 0, ArgPos: 1},
	},
}

func Load(configPath string) Config {
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return defaultConfig
	}

	return cfg
}
