package apiserver

type Config struct {
	BindAddr string `yml:"port`
	LogLevel string `yml:"level`
}

// NewConfig...
func NewConfig(port, level string) *Config {
	return &Config{
		BindAddr: port,
		LogLevel: level,
	}
}
