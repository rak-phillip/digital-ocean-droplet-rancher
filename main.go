package main

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

type doConfig struct {
	dropletName       string
	accessToken       string
	sshFingerprint    string
	url               string
	branch            string
	rancherVersion    string
	bootstrapPassword string
}

func main() {
	var config doConfig

	app := &cli.App{
		Name:    "dorm (Digital Ocean Rancher Manager)",
		Version: "0.0.1",
		Authors: []*cli.Author{
			{
				Name:  "Phillip Rak",
				Email: "rak.phillip@gmail.com",
			},
			{
				Name:  "Jordon Leach",
				Email: "Jordonleach@gmail.com",
			},
		},
		HelpName: "dorm",
		Usage:    "Quickly provision Rancher setups on Digital Ocean",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "droplet-name",
				Usage:       "Name for your Droplet",
				Required:    true,
				Destination: &config.dropletName,
			},
			&cli.StringFlag{
				Name:        "access-token",
				Usage:       "Digital Ocean personal access token",
				Required:    canReadEnv("DO_ENV_ACCESS_TOKEN"),
				Destination: &config.accessToken,
				EnvVars:     []string{"DO_ENV_ACCESS_TOKEN"},
			},
			&cli.StringFlag{
				Name:        "ssh-fingerprint",
				Usage:       "Fingerprint for SSH Public Key",
				Required:    canReadEnv("DO_ENV_SSH_FINGERPRINT"),
				Destination: &config.sshFingerprint,
				EnvVars:     []string{"DO_ENV_SSH_FINGERPRINT"},
			},
			&cli.StringFlag{
				Name:        "url",
				Usage:       "Github url to provision",
				Value:       "https://github.com/rancher/dashboard.git",
				DefaultText: "https://github.com/rancher/dashboard.git",
				Destination: &config.url,
			},
			&cli.StringFlag{
				Name:        "branch",
				Usage:       "Git branch to target",
				Value:       "master",
				DefaultText: "master",
				Destination: &config.branch,
			},
			&cli.StringFlag{
				Name:        "rancher-version",
				Usage:       "Target version of Rancher",
				Value:       "v2.6-head",
				DefaultText: "v2.6-head",
				Destination: &config.rancherVersion,
			},
			&cli.StringFlag{
				Name:        "bootstrap-password",
				Usage:       "Bootstrap password for Rancher",
				Value:       uuid.New().String(),
				Destination: &config.bootstrapPassword,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Provisioning Digital Ocean Droplet...")
			digitalOceanID, ipAddr, _ := createDroplet(&config)

			fmt.Println("Your droplet as been created")
			fmt.Println("DigitalOcean ID: ", digitalOceanID)
			fmt.Println("IP Address: ", ipAddr)
			fmt.Println("Bootstrap Password: ", config.bootstrapPassword)
			return nil
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

}
