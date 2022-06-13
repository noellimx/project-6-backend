package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type GlobalConfig struct {
	OAuth struct {
		Google struct {
			ClientId     string `json:"client_id"`
			ClientSecret string `json:"client_secret"`
		} `json:"google"`
	} `json:"oAuth"`

	Network struct {
		Domain string `json:"domain"`
		Port   string `json:"port"`
	}

	Session struct {
		Key string `json:"key"`
	} `json:"session"`

	Https struct {
		Paths struct {
			CertFileParentVar string `json:"cert_file_parent_var"`
			Certificate       string `json:"certificate"`
			Key               string `json:"key"`
		} `json:"paths"`
	}
}

func ReadConfig(path string) GlobalConfig {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	globalConfig := GlobalConfig{}
	json.Unmarshal(byteValue, &globalConfig)

	return globalConfig
}
