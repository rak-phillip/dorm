package main

import (
	"io/ioutil"
	"strings"
)

// Take env and write to cloud-config
func CloudInit() {
	file := "./cloud-config-tmp"
	d, _ := ioutil.ReadFile(file)

	UpdateConfig(string(d))
}

func UpdateConfig(fileString string) {
	f := fileString
	f = strings.Replace(f, "<REPO_BRANCH>", REPO_BRANCH, -1)
	f = strings.Replace(f, "<REPO_URL>", REPO_URL, 1)
	f = strings.Replace(f, "<RANCHER_VERSION>", RANCHER_VERSION, 1)

	data := []byte(f)
	_ = ioutil.WriteFile("./cloud-config", data, 0o600)
}
