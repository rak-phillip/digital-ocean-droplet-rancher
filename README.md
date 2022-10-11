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

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --droplet-name value        Name for your Droplet
   --access-token value        Digital Ocean personal access token [$DO_ENV_ACCESS_TOKEN]
   --ssh-fingerprint value     Fingerprint for SSH Public Key [$DO_ENV_SSH_FINGERPRINT]
   --url value                 Github url to provision (default: https://github.com/rancher/dashboard.git)
   --branch value              Git branch to target (default: master)
   --rancher-version value     Target version of Rancher (default: v2.7-head)
   --bootstrap-password value  Bootstrap password for Rancher (default: "587ea425-c55b-4783-9ca8-b79880e77636")
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
--rancher-version v2.7-head
```

The build will take around 10 minutes to complete. Once the build is completed you can access your Rancher instance at the provided IP.

### Using Environment Variables  

`dorm` supports assigning environment variables for basic configuration. `dorm` will look for environment variables under `$HOME/.config/dorm/dorm_variables`. Environment variables can also be set in a variety of ways via `~/.profile`, `/etc/profile.d`, `~/.zshrc`, etc...

Supported environment variables:

* `DORM_ENV_ACCESS_TOKEN`: Digital Ocean personal access token 

* `DORM_ENV_SSH_FINGERPRINT`: Fingerprint for your SSH Public Key

### Accessing your instance

SSH into your droplet with the ssh key that matches your fingerprint in DigitalOcean.

```sh
ssh -i root@<droplet-ip> ~/path/to/key
```

You can find the build logs from [cloud-init](https://cloudinit.readthedocs.io/en/latest/) in `/var/log/cloud-init-ouput.log`:

```sh
less /var/log/cloud-init-output.log
```
