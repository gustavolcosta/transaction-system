package configs

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv(pathEnv string) {
	err := godotenv.Load(pathEnv)
	if err != nil {
		log.Fatalf("Error to load enviroments var: %v", err)
	}
}
