//go:build !dev
// +build !dev

package config

func Load() *Config {
	cfg := loadBase()
	cfg.Environment = Prod
	return cfg
}
