package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"proxy/resources"
)

type Config struct {
	Client struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"client"`
	Target struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	ClientToken string `yaml:"clientToken"`
	Keys        struct {
		Incoming string `yaml:"incoming"`
		Outgoing string `yaml:"outgoing"`
	} `yaml:"keys"`
	Log struct {
		Trace   bool `yaml:"trace"`
		Debug   bool `yaml:"debug"`
		UseHex  bool `yaml:"useHex"`
		Packets struct {
			Incoming bool `yaml:"incoming"`
			Outgoing bool `yaml:"outgoing"`
		} `yaml:"packets"`
	} `yaml:"log"`
}

var SavedConfig *Config

// TODO: add 2 other ways to store configs: secure registry or json files

// GetConfig - return the parsed config file
func GetConfig() Config {
	if SavedConfig != nil {
		return *SavedConfig
	}
	conf := readFile()
	SavedConfig = conf
	return *SavedConfig
}

// readFile - try to open the config file and parse into a struct
func readFile() *Config {
	path := resources.GetExecutablePath()
	f, err := os.Open(filepath.Join(path, "config", "config.yaml"))
	if err != nil {
		fmt.Printf("Could not find the config.yaml file: %s\n", err.Error())
		os.Exit(1)
	}
	defer f.Close()

	var config Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	// t
	if err != nil {
		fmt.Printf("Error while reading config.yaml: %s\n", err.Error())
		os.Exit(0)
	}
	return &config
}
