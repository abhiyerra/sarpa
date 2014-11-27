/*

- [ ] Connect to etcd.
- [ ] For each of the services on the configuration.
- [ ] Check and update the configuration file
- [ ] Start or restart the service.
*/
package main

import (
	"fmt"
	"strings"
)

type Nginx struct {
	ServerNames []string
	Proxies     []string
}

func (n *Nginx) Config() string {

	var ProxyPasses []string
	for _, i := range n.Proxies {
		ProxyPasses = append(ProxyPasses, fmt.Sprintf("        proxy_pass %s;", i))
	}

	return fmt.Sprintf(`server {
    server_name %s;
    listen 8080;

    location / {
%s
    }
}`,
		strings.Join(n.ServerNames, ", "),
		strings.Join(ProxyPasses, "\n"),
	)

}

func main() {
	ng := Nginx{
		ServerNames: []string{"a", "b"},
		Proxies:     []string{"b", "c"},
	}

	fmt.Println(ng.Config())
}
