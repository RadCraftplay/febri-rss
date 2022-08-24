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
}
