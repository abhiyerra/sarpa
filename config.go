package main

import (
	"encoding/json"
	"github.com/coreos/go-etcd/etcd"
	"io/ioutil"
	"log"
)

type Config struct {
	Services   []Service    `json:"services"`
	EtcdHosts  []string     `json:"etcd_hosts"`
	etcdClient *etcd.Client `json:"-"`
}

func (c *Config) Parse(configFile string) {
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(file, &c); err != nil {
		log.Fatal(err)
	}
}

func (c *Config) EtcdConnect() {
	c.etcdClient = etcd.NewClient(c.EtcdHosts)

}

func (c *Config) StartWatchmen(restart chan bool) {
	for i := range config.Services {
		go config.Services[i].Watchman(c.etcdClient, restart)
	}
}
