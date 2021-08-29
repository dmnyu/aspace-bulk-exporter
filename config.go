package aspace_bulk_exporter

import (
	"github.com/nyudlts/go-aspace"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func ConfigExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func GetAspaceCreds(configPath string) (map[string]aspace.Creds, error) {
	credsMap := map[string]aspace.Creds{}

	configBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return map[string]aspace.Creds{}, err
	}

	err = yaml.Unmarshal(configBytes, &credsMap); if err != nil {
		return map[string]aspace.Creds{}, err
	}

	return credsMap, nil

}

func GetAspaceEnvironments(creds map[string]aspace.Creds) []string {
	environments := []string{}
	for k,_ := range creds {
		environments = append(environments, k)
	}
	return environments
}