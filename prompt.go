package main

import (
	"fmt"
)

var DROPLET_NAME, DIGITAL_OCEAN_ACCESS_TOKEN, SSH_FINGERPRINT string
var REPO_URL, REPO_BRANCH string

func Prompt() (string, string, string, string, string) {
	fmt.Println("Give a name for your droplet:")
	fmt.Scanln(&DROPLET_NAME)

	fmt.Println("Your forked repo:")
	fmt.Scanln(&REPO_URL)

	fmt.Println("The branch to build from:")
	fmt.Scanln(&REPO_BRANCH)

	fmt.Println("DigitalOcean access token:")
	fmt.Scanln(&DIGITAL_OCEAN_ACCESS_TOKEN)

	fmt.Println("ssh key fingerprint:")
	fmt.Scanln(&SSH_FINGERPRINT)

	return DROPLET_NAME, DIGITAL_OCEAN_ACCESS_TOKEN, SSH_FINGERPRINT, REPO_URL, REPO_BRANCH
}
