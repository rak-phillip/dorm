package main

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/digitalocean/godo"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

func CreateDroplet(config *Config) (int, error) {
	client := godo.NewFromToken(config.accessToken)
	allKeys := []godo.DropletCreateSSHKey{
		{Fingerprint: config.sshFingerprint},
	}

	// tags := []string{"tag-test"}
	CloudInit(config)
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
		return 0, err
	}

	ipAddr := WaitForDroplet(ctx, client, newDroplet.ID)

	ConnectToHost(ipAddr)

	return newDroplet.ID, nil
}

func WaitForDroplet(ctx context.Context, client *godo.Client, dropletId int) string {
	var ipAddr string
	for {
		droplet, _, _ := client.Droplets.Get(ctx, dropletId)
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

func ConnectToHost(host string) (*ssh.Client, *ssh.Session, error) {
	fmt.Println("Enter a password: ")
	pass, _ := term.ReadPassword(int(syscall.Stdin))

	sshConfig := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password(string(pass)),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}
