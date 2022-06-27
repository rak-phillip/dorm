# Development Droplet for Rancher Dashboard

This Go program will create a testable instance of [Rancher](https://github.com/rancher/rancher) with a custom [dashboard](https://github.com/rancher/dashboard), hosted on DigitalOcean.

## Prerequisites

1. A fork of the [dashboard](https://github.com/rancher/dashboard) repo
2. DigitalOcean AccessToken
3. SSH fingerprint - this can be found within DigitalOcean settings

## Usage

```
NAME:
   dorm (Digital Ocean Rancher Manager) - Quickly provision Rancher setups on Digital Ocean

USAGE:
   dorm [global options] command [command options] [arguments...]

GLOBAL OPTIONS:
   --droplet-name value        Name for your Droplet
   --access-token value        Digital Ocean personal access token (default: "c63a4d22e0724ebfc566178c8a858c7f9451240b4b41372a6e0fbaabe4d2bfb8") [$DO_ENV_ACCESS_TOKEN]
   --ssh-fingerprint value     Fingerprint for SSH Public Key (default: "8c:96:3b:ea:4b:bc:5c:54:2b:11:20:9e:d4:30:3f:c2") [$DO_ENV_SSH_FINGERPRINT]
   --url value                 Github url to provision (default: https://github.com/rancher/dashboard.git)
   --branch value              Git branch to target (default: master)
   --rancher-version value     Target version of Rancher (default: v2.6-head)
   --bootstrap-password value  Bootstrap password for Rancher (default: "d6538231-c73a-4262-a598-5a2fc09dff58")
   --help, -h                  show help (default: false)
   --version, -v               print the version (default: false)
```
### Installing dorm

`dorm` requires [a supported release of Go](https://go.dev/doc/devel/release#policy).

```
$ go install github.com/rak-phillip/dorm@latest
```

To find out where `dorm` was installed you can run `go list -f {{.Target}} github.com/rak-phillip/dorm`. For `dorm` to be used globally add that directory to the `$PATH` environment setting.

### Running the program

```sh
$ ./dorm --droplet-name my-first-rancher-droplet \
--access-token { your-digital-ocean-access-token } \
--ssh-fingerprint { your-ssh-public-key-fingerprint } \
--url https://github.com/rak-phillip/dashboard.git \
--branch master \
--rancher-version v2.6-head
```

The build will take around 10 minutes to complete. Once the build is completed you can access your Rancher instance at the provided IP.

### Accessing your instance

SSH into your droplet with the ssh key that matches your fingerprint in DigitalOcean.

```sh
ssh -i root@<droplet-ip> ~/path/to/key
```

You can find the build logs from [cloud-init](https://cloudinit.readthedocs.io/en/latest/) in `/var/log/cloud-init-ouput.log`:

```sh
less /var/log/cloud-init-output.log
```
