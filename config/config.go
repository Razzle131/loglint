package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

// Func describes a function to analyze, e.g. [slog.Info].
type Func struct {
	Name    string `yaml:"name"`   // full function name (example: (*log/slog.Logger).Info)
	MsgPos  int    `yaml:"msgPos"` // position of msg argument starting from 0
	ArgsPos int    `yaml:"argPos"` // The position of the "args ...any" argument starting from 0
}

type Config struct {
	EnabledRules int      `yaml:"enabled_rules" env:"ENABLED_RULES"` // bit mask for rules on/off
	AvoidedData  []string `yaml:"avoided_data" env:"AVOIDED_DATA"`   // slice of avoided names, lowercase only
	EnabledFuncs []Func   `yaml:"enabled_funcs"`                     // slice of func definitions that linter searches for
}

func MustLoad(configPath string) Config {
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config %q: %s", configPath, err)
	}

	return cfg
}
