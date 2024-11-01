package apiserver

// Config ...
type Config struct {
	BinAddr string `toml:"bind_addr`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BinAddr: "8080",
	}
}
