package config

type DatabaseConfig struct {
	Path string
}

func LoadDatabaseConfig(envs map[string]string) DatabaseConfig {

	return DatabaseConfig{
		Path: envs["DB_PATH"],
	}
}
