package config

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	configFileName = "moodle-api.yml"
)

type MoodleApiConfig struct {
	URL   string `yaml:"url"`
	Token string `yaml:"token"`
}

// ReadConfigFile read configuration file
func readConfigFile() (cfg *MoodleApiConfig) {

	f, err := os.Open(configFileName)
	if err != nil {
		d, _ := yaml.Marshal(&MoodleApiConfig{})
		_ = ioutil.WriteFile(configFileName, d, 0644)
		log.Fatalln(err.Error(), "- creating config.yml")
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	return cfg
}

func NewMoodleApiConfig() (*MoodleApiConfig, error) {

	config := readConfigFile()

	if config.URL != "" && config.Token != "" {
		if strings.HasSuffix(config.URL, "/") {
			config.URL = config.URL[:len(config.URL)-1]
		}

		return config, nil
	}

	return nil, errors.New("domain or token null")
}
