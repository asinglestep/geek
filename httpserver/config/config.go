package config

type Config struct {
	Server ServerConfig `toml:"server"`
}

type ServerConfig struct {
	Address string `toml:"address"`
}
