//go:build dev
// +build dev

package config

func Load() *Config {
	cfg := loadBase()
	cfg.Environment = Dev
	return cfg
}
