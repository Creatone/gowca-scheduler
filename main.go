package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"intel.com/gowca-scheduler/pkg/config"
	"intel.com/gowca-scheduler/pkg/server"
	"io/ioutil"
)

var configPath = flag.String("config", "examples/gowca_scheduler_config.yaml", "Path to configuration file.")

func main() {
	flag.Parse()

	configBytes, err := ioutil.ReadFile(*configPath)
	if err != nil {
		fmt.Println(err)
	}

	var config config.Config

	// Parse it.
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		fmt.Println(err)
	}

	server := server.Server{Host: config.Host, Port: config.Port}
	server.Serve()
}
