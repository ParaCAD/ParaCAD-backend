package utils

type Config struct {
	Port string
}

func MustLoadConfig() *Config {
	envs := mustGetFromEnvs(
		"PORT",
	)

	return &Config{
		Port: envs["PORT"],
	}
}
