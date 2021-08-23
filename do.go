package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/digitalocean/godo"
)

func main() {
	fmt.Println("Hello, World!")
	digitalOceanId, _ := CreateDroplet()
	fmt.Println(digitalOceanId)
}

func CreateDroplet() (int, error) {
	client := godo.NewFromToken(os.Getenv("DIGITAL_OCEAN_ACCESS_TOKEN"))

	dropletName := "rancher.prak"
	tags := []string{"prak"}

	createRequest := &godo.DropletCreateRequest{
		Name:   dropletName,
		Region: "sfo3",
		Size:   "s-2vcpu-4gb",
		Tags:   tags,
		Image: godo.DropletCreateImage{
			Slug: "ubuntu-20-04-x64",
		},
	}

	ctx := context.TODO()

	newDroplet, _, err := client.Droplets.Create(ctx, createRequest)

	if err != nil {
		fmt.Printf("Something bad happened: %s\n\n", err)
		return 0, err
	}

	WaitForDroplet(ctx, client, newDroplet.ID)

	return newDroplet.ID, nil
}

func WaitForDroplet(ctx context.Context, client *godo.Client, dropletId int) {
	for {
		droplet, _, _ := client.Droplets.Get(ctx, dropletId)
		status := droplet.Status
		if status == "active" {
			fmt.Println("Node is ready.")
			break
		} else {
			fmt.Println("Node is not ready.")
			time.Sleep(8 * time.Second)
		}
	}
}
