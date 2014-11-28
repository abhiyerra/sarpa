package main

import (
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"log"
	"strings"
	"time"
)

type Service struct {
	Name        string   `json:"service_name"`
	Hosts       []string `json:"hosts"`
	Ssl         bool     `json:"ssl"`
	CertPath    string   `json:"cert_path"`
	ProxyPasses []string `json:"-"`
}

func (n *Service) Proxies() string {
	var passes []string

	for _, i := range n.ProxyPasses {
		passes = append(passes, fmt.Sprintf("        proxy_pass %s;", i))
	}

	return strings.Join(passes, "\n")
}

func (n *Service) HostNames() string {
	return strings.Join(n.Hosts, ", ")
}

func (n *Service) Config() string {
	return fmt.Sprintf(`server {
    server_name %s;
    listen 8080;

    location / {
%s
    }
}`,
		n.HostNames(),
		n.Proxies(),
	)
}

func etcdServicePath(serviceName string) string {
	return fmt.Sprintf("/sarpa/%s", serviceName)
}

func (n *Service) Watchman(client *etcd.Client) {
	log.Println("Starting watchman for", n.Name)

	servicePath := etcdServicePath(n.Name)

	for {
		// Get the nodes and their values.
		resp, err := client.Get(servicePath, false, false)
		if err == nil {
			var newProxies []string
			for _, n := range resp.Node.Nodes {
				log.Printf("%s: %s\n", n.Key, n.Value)
				newProxies = append(newProxies, n.Value)
			}
			n.ProxyPasses = newProxies

		} else {
			log.Println(err)
		}

		log.Println(n.Config())

		// Watch for changes to the values
		watchChan := make(chan *etcd.Response)
		go client.Watch(servicePath, 0, true, watchChan, nil)
		log.Println("Waiting for an update...")

		select {
		case r := <-watchChan:
			log.Printf("Got updated creds: %s: %s\n", r.Node.Key, r.Node.Value)
		case <-time.After(time.Second * 10):
			log.Println(n.Name, "timeout. Watching again")
		}
	}
}
