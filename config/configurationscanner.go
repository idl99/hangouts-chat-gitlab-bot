package config

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Configuration - struct to represent the application configuration
type Configuration struct {
	Chat struct {
		WebhookURL string `yaml:"webhookURL"`
	}
}

func processError(err error) {
	log.Fatal(err)
	os.Exit(2)
}

// ReadFile - this function will read the application configuration from the properties file
func ReadFile(cfg *Configuration) {
	data, err := ioutil.ReadFile("config/properties.yml")
	if err != nil {
		processError(err)
	}

	decoder := yaml.NewDecoder(bytes.NewBuffer(data))
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}
