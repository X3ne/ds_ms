package config

type Config struct {
	LogLevel	string
	SERVER		ServerConfig
	DB				DatabaseConfig
}

func NewConfig() *Config {
	envs := ValidateEnvs()

	return &Config{
		LogLevel: envs["LOG_LEVEL"],
		SERVER: LoadServerConfig(envs),
		DB: LoadDatabaseConfig(envs),
	}
}
