package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var dropletName string
	var accessToken string
	var sshFingerprint string
	var url string
	var branch string
	var rancherVersion string

	app := &cli.App{
		Name:  "Rancher Digital Ocean Provisioner",
		Usage: "this is a nice usage instruction",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "droplet-name",
				Usage:       "Name for your Droplet",
				Required:    true,
				Destination: &dropletName,
			},
			&cli.StringFlag{
				Name:        "access-token",
				Usage:       "Digital Ocean personal access token",
				Required:    true,
				Destination: &accessToken,
			},
			&cli.StringFlag{
				Name:        "ssh-fingerprint",
				Usage:       "Fingerprint for SSH Public Key",
				Required:    true,
				Destination: &sshFingerprint,
			},
			&cli.StringFlag{
				Name:        "url",
				Usage:       "Github url for project",
				Required:    true,
				Destination: &url,
			},
			&cli.StringFlag{
				Name:        "branch",
				Usage:       "Git branch provisioning target",
				DefaultText: "master",
				Destination: &branch,
			},
			&cli.StringFlag{
				Name:        "rancher-version",
				Usage:       "Target version of Rancher",
				DefaultText: "v2.6-head",
				Destination: &rancherVersion,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Provisioning Digital Ocean Droplet...")
			digitalOceanId, _ := CreateDroplet(dropletName, accessToken, sshFingerprint, url, branch, rancherVersion)

			fmt.Println("Your droplet as been created")
			fmt.Println("DigitalOcean ID: ", digitalOceanId)
			return nil
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}
