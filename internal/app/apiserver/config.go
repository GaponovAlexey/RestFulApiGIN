package apiserver

type Config struct {
	BindAddr string `toml:"bind_addr`
}


// NewConfig...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":3000",
	}
}
