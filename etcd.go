package main

import (
	"fmt"
	"github.com/coreos/go-etcd/etcd"
)

func EtcdServicePath(serviceName string) string {
	return fmt.Sprintf("/sarpa/%s", serviceName)
}

// TODO: For client use.
func EtcdInsertHost(client *etcd.Client, serviceName, host string) bool {
	// Create a directory in /sarpa/:service_name
	// Create a random child in the directory with key=host value=ip. Set ttl=30seconds
	// On quit kill the child if possible
	return false
}
