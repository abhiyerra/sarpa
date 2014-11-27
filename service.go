package main

import (
	"fmt"
	"strings"
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
