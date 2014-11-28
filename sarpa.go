package main

import (
	"flag"
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
	restart := make(chan bool)

	config.StartWatchmen(restart)

	for {
		log.Println("restarting")
		<-restart
	}
}
