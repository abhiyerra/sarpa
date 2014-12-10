package main

import (
	"github.com/coreos/go-etcd/etcd"
	"log"
	"testing"
)

var (
	client *etcd.Client
)

func SetupEtcd() {
	// Setup
	client = etcd.NewClient([]string{"http://127.0.0.1:4001"})
	client.CreateDir("/sarpa", 200)

	client.CreateDir("/sarpa/blah", 200)
	client.Create("/sarpa/blah/asd", "1", 200)
	client.Create("/sarpa/blah/meh", "2", 200)

	client.CreateDir("/sarpa/hay", 200)
	client.AddChild("/sarpa/hay", "4", 200)
	client.AddChild("/sarpa/hay", "5", 200)
	client.AddChild("/sarpa/hay", "6", 200)
}

func CleanUpEtcd() {
	// Clean up
	client.Delete("/sarpa", true)
}

func TestSarpa(t *testing.T) {
	SetupEtcd()

	config := &Config{}
	config.EtcdConnect([]string{"http://127.0.0.1:4001"})
	config.AwsConnect()

	config.GetServices()

	if len(config.Services["blah"]) != 2 {
		t.Error("blah doesn't have enough nodes")
	}

	if len(config.Services["hay"]) != 3 {
		t.Error("hay doesn't have enough nodes")
	}

	log.Println(string(config.jsonedServices()))

	CleanUpEtcd()
}
