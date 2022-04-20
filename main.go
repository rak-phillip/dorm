package main

import (
	"fmt"
)

func main() {
	fmt.Println("Provisioning Digital Ocean Droplet...")
	digitalOceanId, _ := CreateDroplet(Prompt())

	fmt.Println("Your droplet as been created")
	fmt.Println("DigitalOcean ID: ", digitalOceanId)
}
