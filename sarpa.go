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

type Service struct {
	Name     string   `json:"service_name"`
	Hosts    []string `json:"hosts"`
	Ssl      bool     `json:"ssl"`
	CertPath string   `json:"cert_path"`
}

// type Nginx struct {
// 	ServerNames []string
// 	Proxies     []string
// }

// func (n *Service) Config() string {

// 	var ProxyPasses []string
// 	for _, i := range n.Proxies {
// 		ProxyPasses = append(ProxyPasses, fmt.Sprintf("        proxy_pass %s;", i))
// 	}

// 	return fmt.Sprintf(`server {
//     server_name %s;
//     listen 8080;

//     location / {
// %s
//     }
// }`,
// 		strings.Join(n.ServerNames, ", "),
// 		strings.Join(ProxyPasses, "\n"),
// 	)

// }

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
