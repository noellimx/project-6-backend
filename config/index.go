package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type JWT struct {
	Secret string `json:"secret"`
}
type PSQL struct {
	Username     string `json:"username"`
	DatabaseName string `json:"database_name"`
	Host         string `json:"host"`
	Port         string `json:"port"`
}

type GlobalConfig struct {
	JWT   JWT  `json:"jwt"`
	PSQL  PSQL `json:"psql"`
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

func readConfig(path string) *GlobalConfig {
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

	return &globalConfig
}

type Environment int

const (
	Production Environment = iota
	Test
)

func ReadConfig(env Environment) *GlobalConfig {

	configFileParent := os.Getenv("HOME")

	var subpath string

	if env == Production {
		subpath = "production"
	} else if env == Test {
		subpath = "test"
	} else {
		log.Fatal("Environment not supported")
	}

	configFilePath := configFileParent + "/customkeystore/" + subpath + "/config.json"

	return readConfig(configFilePath)
}
