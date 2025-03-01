package utils

import (
	"log"
	"os"
)

type Config struct {
	Port       string
	JWTSecret  string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
}

func MustLoadConfig() *Config {
	envs := mustGetFromEnvs(
		"PORT",
		"JWT_SECRET",
		"DB_HOST",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
	)

	return &Config{
		Port:       envs["PORT"],
		JWTSecret:  envs["JWT_SECRET"],
		DBHost:     envs["DB_HOST"],
		DBUser:     envs["DB_USER"],
		DBPassword: envs["DB_PASSWORD"],
		DBName:     envs["DB_NAME"],
	}
}

func mustGetFromEnvs(keys ...string) map[string]string {
	PrintLine("Envs")
	defer PrintLine("")
	envs := map[string]string{}
	for _, key := range keys {
		envs[key] = os.Getenv(key)
		if envs[key] != "" {
			log.Printf("%s: %s\n", green(key), green(envs[key]))
		} else {
			log.Fatalf("%s environment variable not set!", red(key))
		}
	}
	return envs
}
