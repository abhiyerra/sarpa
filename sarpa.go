package main

import (
	"flag"
	"fmt"
	"log"
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
	config.EtcdConnect()
}

func main() {
	config.StartWatchmen()

	var i int
	fmt.Scanf("%d", &i)
}
