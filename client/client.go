// Sarpa client which updates etcd with the correct host ip.
package sarpa

import (
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"log"
)

func etcdServicePath(serviceName string) string {
	return fmt.Sprintf("/sarpa/%s", serviceName)
}

// This method should be called every ttl to make sure that etcd has
// the current value and sarpa can still talk to this host.
//
//  - client: The etcd client used to talk to etcd.
//  - serviceName: The name of the service which will be proxied to from sarpa.
//  - hostIp: Update the hostIp so that sarpa can connect to this host.
//  - ttl: How long to keep this host alive for.
//
// This method should be run every ttl seconds. So that sarpa can have
// the correct hosts.
func SarpaUpdater(client *etcd.Client, serviceName, hostIp string, ttl int) bool {
	servicePath := etcdServicePath(serviceName)

	if _, err := client.Get(servicePath, false, false); err != nil {
		// Directory doesn't exist.

		_, err := client.CreateDir(servicePath, 300)
		if err != nil {
			log.Println(err)
		}
	}

	// Create a random child in the directory with value=host. Set ttl=30seconds
	_, err := client.AddChild(servicePath, hostIp, 30)
	if err != nil {
		log.Println(err)
	}

	return false
}
