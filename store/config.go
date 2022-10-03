package store

type Config struct {
	URI string
}

func NewConfig() *Config {
	return &Config{}
}