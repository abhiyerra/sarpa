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
	// ng := Nginx{
	// 	ServerNames: []string{"a", "b"},
	// 	Proxies:     []string{"b", "c"},
	// }

	// fmt.Println(ng.Config())

}
