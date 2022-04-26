# Development Droplet for Rancher Dashboard

This Go program will create a testable instance of [Rancher](https://github.com/rancher/rancher) with a custom [dashboard](https://github.com/rancher/dashboard), hosted on DigitalOcean.

## Prerequisites

1. A fork of the [dashboard](https://github.com/rancher/dashboard) repo
2. DigitalOcean AccessToken
3. SSH fingerprint - this can be found within DigitalOcean settings

## How to use

You can run the program with:

```sh
go run .
```

After a few questions the droplet will be created and you'll be prompted to give a root password for your droplet. The build will take around 10 minutes to complete. Once the build is completed you can access your Rancher instance at the provided IP.

> _Note_: Currently you will need to ssh into your droplet and retrieve the Bootstrap Password - anyone is welcome to fix this...

### Accessing your instance

SSH into your droplet with the ssh key that matches your fingerprint in DigitalOcean.

```sh
ssh -i root@<droplet-ip> ~/path/to/key
```

You can find the build logs from [cloud-init](https://cloudinit.readthedocs.io/en/latest/) in `/var/log/cloud-init-ouput.log`:

```sh
less /var/log/cloud-init-output.log
```
