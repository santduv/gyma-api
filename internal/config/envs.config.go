package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Environments struct {
	MongoURI string `env:"MONGO_URI,required"`
	MongoDB  string `env:"MONGO_DB,required"`
}

var Envs = getEnvs()

func getEnvs() *Environments {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file : %v", err)
	}

	envs := &Environments{}
	err = env.Parse(envs)

	if err != nil {
		log.Fatalf("Error parsing env vars %v", err)
	}

	return envs
}
