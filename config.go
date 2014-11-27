package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Services []Service
}

func (c *Config) Parse(configFile string) {
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(file, &c.Services); err != nil {
		log.Fatal(err)
	}
}
