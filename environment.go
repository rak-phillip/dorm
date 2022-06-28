package main

import (
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
)

func canReadEnv(envVar string) bool {
	godotenv.Load(getUsrHome())
	err := godotenv.Load()

	accessToken := os.Getenv(envVar)

	return err != nil || accessToken == ""
}

func getUsrHome() string {
	usr, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	return path.Join(usr, ".config/dorm/dorm_variables")
}
