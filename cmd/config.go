package cmd

import (
	"github.com/spf13/viper"
)

// Log config struct
type Log struct {
	// Level is the minimum log level that should appear on the output.
	Level string
}

// Config is the namespacex config file
type Config struct {
	Log
}

// Configure configures some defaults in the Viper instance.
func Configure(v *viper.Viper) {
	v.SetDefault("log.level", "info")
	v.SetDefault("log.formatter.name", "json")
	v.AllowEmptyEnv(true)
}
