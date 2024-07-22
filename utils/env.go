package utils

import (
	"log"
	"os"
)

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
