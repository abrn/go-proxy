package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Client struct{
		Host	string 	`yaml:"host"`
		Port	int		`yaml:"port"`
	} `yaml:"client"`
	Target struct{
		Host	string 	`yaml:"host"`
		Port	int		`yaml:"port"`
	} `yaml:"server"`
	Keys struct{
		Incoming	string `yaml:"incoming"`
		Outgoing	string `yaml:"outgoing"`
	} `yaml:"keys"`
	Log struct{
		Trace		bool `yaml:"trace"`
		Debug 		bool `yaml:"debug"`
		Packets struct{
			Incoming	bool `yaml:"incoming"`
			Outgoing 	bool `yaml:"outgoing"`
		} `yaml:"packets"`
	} `yaml:"log"`
}

// GetConfig - return the parsed config file
func GetConfig() Config {
	conf := readFile()
	return *conf
}

// readFile - try to open the config file and parse into a struct
func readFile() *Config {
	f, err := os.Open("config.yaml")
	if err != nil {
		fmt.Printf("Error while opening config.yaml: %s\n", err.Error())
		os.Exit(0)
	}
	defer f.Close()

	var config Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Printf("Error while parsing config.yaml: %s\n", err.Error())
		os.Exit(0)
	}
	return &config
}