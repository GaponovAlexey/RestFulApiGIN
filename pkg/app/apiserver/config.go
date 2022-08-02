package apiserver

import "github.com/gaponovalexey/go-restapi/pkg/app/store"

type Config struct {
	BindAddr string `yml:"port`
	LogLevel string `yml:"level`
	Store    *store.Config
}

// NewConfig...
func NewConfig(port, level, db string) *Config {
	return &Config{
		BindAddr: port,
		LogLevel: level,
		Store: db,
	}
}
