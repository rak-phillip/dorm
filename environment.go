package main

import "github.com/joho/godotenv"

func init() {
	godotenv.Load()
}

func canReadEnv(envVar string) bool {
	var doEnv map[string]string
	doEnv, envErr := godotenv.Read()
	accessToken := doEnv[envVar]

	return envErr != nil || accessToken == ""
}
