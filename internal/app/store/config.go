package store

// Config ...
type Config struct {
	DatabaseHost     string `env:"DB_HOST"`
	DatabasePort     string `env:"DB_PORT"`
	DatabaseUser     string `env:"DB_USER"`
	DatabasePassword string `env:"DB_PASSWORD"`
	DatabaseName     string `env:"DB_NAME"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
