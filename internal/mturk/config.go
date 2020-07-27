package mturk

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config is MTurk configuration
type Config struct {
	Sandbox bool `mapstructure:"sandbox"`
}

// Validate configuration is ok
func (c *Config) Validate() error {
	return nil
}

// ConfigFlags helps configure cobra and viper flags.
func ConfigFlags(cmd *cobra.Command, prefix string) error {
	if prefix == "" {
		prefix = "mturk"
	}

	viper.SetDefault(prefix, &Config{})

	flag := prefix + ".sandbox"
	val := false
	cmd.Flags().Bool(flag, val, "Use the MTurk Sandbox")
	viper.SetDefault(flag, val)

	return nil
}