# Development Droplet for Rancher Dashboard

This Go program will create a testable instance of [Rancher](https://github.com/rancher/rancher) with a custom [dashboard](https://github.com/rancher/dashboard), hosted on DigitalOcean.

## Prerequisites

1. A fork of the [dashboard](https://github.com/rancher/dashboard) repo
2. DigitalOcean AccessToken
3. SSH fingerprint - this can be found within DigitalOcean settings

## Usage

```
NAME:
   do - Quickly provision Rancher setups on Digital Ocean

USAGE:
   do [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --droplet-name value     Name for your Droplet
   --access-token value     Digital Ocean personal access token
   --ssh-fingerprint value  Fingerprint for SSH Public Key
   --url value              Github url to provision (default: https://github.com/rancher/dashboard.git)
   --branch value           Git branch to target (default: master)
   --rancher-version value  Target version of Rancher (default: v2.6-head)
   --help, -h               show help (default: false)
```
### Running the program

```sh
$ go build
$ ./do --droplet-name my-first-rancher-droplet \
--access-token { your-digital-ocean-access-token } \
--ssh-fingerprint { your-ssh-public-key-fingerprint } \
--url https://github.com/rak-phillip/dashboard.git \
--branch master \
--rancher-version v2.6-head
```

The droplet will be created and you'll be prompted to give a root password for your droplet. The build will take around 10 minutes to complete. Once the build is completed you can access your Rancher instance at the provided IP.

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
