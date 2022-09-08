package common

import (
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	configFilename = "./config.yaml"
)

func getDefaultConfiguration() FebriRssConfiguration {
	return FebriRssConfiguration{
		struct {
			FetchRss struct {
				Period time.Duration "yaml:\"period\""
			}
		}{
			FetchRss: struct {
				Period time.Duration "yaml:\"period\""
			}{
				Period: time.Hour * 6,
			},
		},
		struct {
			AppKey    string "yaml:\"appKey\""
			AppSecret string "yaml:\"appSecret\""
		}{
			AppKey:    "<YOUR APP KEY>",
			AppSecret: "<YOUR APP SECRET>",
		},
		struct {
			Host     string "yaml:\"host\""
			Port     int    "yaml:\"port\""
			Username string "yaml:\"username\""
			Password string "yaml:\"password\""
			DbName   string "yaml:\"dbName\""
		}{
			Host:     "localhost",
			Port:     5432,
			Username: "<YOUR DATABASE USERNAME>",
			Password: "<YOUR DATABASE PASSWORD>",
			DbName:   "<YOUR DATABASE NAME>",
		},
	}
}

func saveConfiguration(configuration FebriRssConfiguration) {
	out, err := yaml.Marshal(&configuration)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(configFilename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	n, err := file.Write(out)
	if err != nil || n != len(out) {
		panic("Unable to properly write configuration!")
	}
}

func LoadConfiguration() FebriRssConfiguration {
	var configuration FebriRssConfiguration

	if _, err := os.Stat(configFilename); os.IsNotExist(err) {
		saveConfiguration(getDefaultConfiguration())
	}

	content, err := ioutil.ReadFile(configFilename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(content, &configuration)
	if err != nil {
		panic(err)
	}
	return configuration
}

type FebriRssConfiguration struct {
	Services struct {
		FetchRss struct {
			Period time.Duration `yaml:"period"`
		}
	}
	FebriApi struct {
		AppKey    string `yaml:"appKey"`
		AppSecret string `yaml:"appSecret"`
	}
	Postgres struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbName"`
	}
}
