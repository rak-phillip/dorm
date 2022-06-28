package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Take env and write to cloud-config
func cloudInit(config *doConfig) {
	var file string = "./cloud-config-tmp"

	if config.useMinConfig == true {
		fmt.Println("Using config ./cloud-config-min")
		file = "./cloud-config-min"
	}

	d, _ := ioutil.ReadFile(file)

	updateDoConfig(string(d), config)
}

func updateDoConfig(fileString string, config *doConfig) {
	f := fileString
	f = strings.Replace(f, "<REPO_BRANCH>", config.branch, -1)
	f = strings.Replace(f, "<REPO_URL>", config.url, 1)
	f = strings.Replace(f, "<RANCHER_VERSION>", config.rancherVersion, 1)
	f = strings.Replace(f, "<RANCHER_BOOTSTRAP_PASSWORD>", config.bootstrapPassword, 1)

	data := []byte(f)
	_ = ioutil.WriteFile("./cloud-config", data, 0o600)
}
