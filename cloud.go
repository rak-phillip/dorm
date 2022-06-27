package main

import (
	"io/ioutil"
	"strings"
)

// Take env and write to cloud-config
func CloudInit(config *Config) {
	file := "./cloud-config-tmp"
	d, _ := ioutil.ReadFile(file)

	UpdateConfig(string(d), config)
}

func UpdateConfig(fileString string, config *Config) {
	f := fileString
	f = strings.Replace(f, "<REPO_BRANCH>", config.branch, -1)
	f = strings.Replace(f, "<REPO_URL>", config.url, 1)
	f = strings.Replace(f, "<RANCHER_VERSION>", config.rancherVersion, 1)
	f = strings.Replace(f, "<RANCHER_BOOTSTRAP_PASSWORD>", config.bootstrapPassword, 1)

	data := []byte(f)
	_ = ioutil.WriteFile("./cloud-config", data, 0o600)
}
