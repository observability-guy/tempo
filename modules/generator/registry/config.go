package registry

import (
	"flag"
	"time"
)

type Config struct {
	// CollectionInterval controls how often to collect metrics.
	// Defaults to 15s.
	CollectionInterval time.Duration `yaml:"collection_interval"`

	// StaleDuration controls how quickly series become stale and are deleted from the registry. An active
	// series is deleted if it hasn't been updated more stale duration.
	// Defaults to 15m.
	StaleDuration time.Duration `yaml:"stale_duration"`

	// ExternalLabels are added to any time series generated by this instance.
	ExternalLabels map[string]string `yaml:"external_labels,omitempty"`
}

// RegisterFlagsAndApplyDefaults registers the flags.
func (cfg *Config) RegisterFlagsAndApplyDefaults(prefix string, f *flag.FlagSet) {
	cfg.CollectionInterval = 15 * time.Second
	cfg.StaleDuration = 15 * time.Minute
}