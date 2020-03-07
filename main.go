package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var configPath = flag.String("config", "/etc/gowca_scheduler_config.yaml", "Path to configuration file.")

func main() {
	flag.Parse()

	// Read config file.
	configBytes, err := ioutil.ReadFile(*configPath)
	if err != nil {
		fmt.Println(err)
	}

	var config Config

	// Parse it.
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		fmt.Println(err)
	}

	// Validate it.
	fmt.Println("Config: ", config)

