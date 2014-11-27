package main

import (
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"log"
)

func EtcdServicePath(serviceName string) string {
	return fmt.Sprintf("/sarpa/%s", serviceName)
}

// TODO: For client use.
func SarpaUpdater(client *etcd.Client, serviceName, host string, ttl int) bool {
	servicePath := EtcdServicePath(serviceName)

	if _, err := client.Get(servicePath, false, false); err != nil {
		// Directory doesn't exist.

		_, err := client.CreateDir(servicePath, 300)
		if err != nil {
			log.Println(err)
		}
	}

	// Create a random child in the directory with value=host. Set ttl=30seconds
	_, err := client.AddChild(servicePath, host, 30)
	if err != nil {
		log.Println(err)
	}

	return false
}
