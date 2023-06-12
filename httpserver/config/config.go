package config

type Config struct {
	Server ServerConfig `toml:"server"`
	Log    LoggerConfig `toml:"log"`
}

type ServerConfig struct {
	Address string `toml:"address"`
}

type LoggerConfig struct {
	Level string `toml:"level"`
}
