package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/digitalocean/godo"
)

func createDroplet(config *doConfig) (int, string, error) {
	client := godo.NewFromToken(config.accessToken)
	allKeys := []godo.DropletCreateSSHKey{
		{Fingerprint: config.sshFingerprint},
	}

	// tags := []string{"tag-test"}
	cloudInit(config)
	data, _ := os.ReadFile("./cloud-config")

	createRequest := &godo.DropletCreateRequest{
		Name:   config.dropletName,
		Region: "sfo3",
		Size:   "s-4vcpu-8gb",
		// Tags:    tags,
		SSHKeys: allKeys,
		Image: godo.DropletCreateImage{
			Slug: "ubuntu-20-04-x64",
		},
		UserData: string(data),
	}

	ctx := context.TODO()

	newDroplet, _, err := client.Droplets.Create(ctx, createRequest)

	if err != nil {
		fmt.Printf("Error: %s\n\n", err)
		return 0, "", err
	}

	ipAddr := waitForDroplet(ctx, client, newDroplet.ID)

	return newDroplet.ID, ipAddr, nil
}

func waitForDroplet(ctx context.Context, client *godo.Client, dropletID int) string {
	var ipAddr string
	for {
		droplet, _, _ := client.Droplets.Get(ctx, dropletID)
		status := droplet.Status
		if status == "active" {
			ip, _ := droplet.PublicIPv4()
			fmt.Printf("%s is ready: %s\n", droplet.Name, ip)
			ipAddr = ip
			break
		} else {
			fmt.Printf("%s is not ready.\n", droplet.Name)
			time.Sleep(8 * time.Second)
		}
	}

	return ipAddr
}
