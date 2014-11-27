package main

import (
	"flag"
	"log"
	//	"fmt"
	//	"strings"
)

var (
	config *Config
)

func init() {
	var configFile = flag.String("config", "", "configuration for services")
	flag.Parse()

	if *configFile == "" {
		log.Fatal("Need config file")
	}

	config = &Config{}
	config.Parse(*configFile)
}

func main() {
	for i := range config.Services {
		log.Println(config.Services[i].Config())
	}
}
