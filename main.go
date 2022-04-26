package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Rancher Digital Ocean Provisioner",
		Usage: "this is a nice usage instruction",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "droplet-name",
				Usage: "Name for your Droplet",
			},
			&cli.StringFlag{
				Name:  "url",
				Usage: "Github url for project",
			},
			&cli.StringFlag{
				Name:  "branch",
				Usage: "Git branch provisioning target",
			},
			&cli.StringFlag{
				Name:  "rancher-version",
				Usage: "Target version of Rancher",
			},
			&cli.StringFlag{
				Name:  "access-token",
				Usage: "Digital Ocean personal access token",
			},
			&cli.StringFlag{
				Name:  "ssh-fingerprint",
				Usage: "Fingerprint for SSH Public Key",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Some Action")
			return nil
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Provisioning Digital Ocean Droplet...")
	digitalOceanId, _ := CreateDroplet(Prompt())

	fmt.Println("Your droplet as been created")
	fmt.Println("DigitalOcean ID: ", digitalOceanId)
}
