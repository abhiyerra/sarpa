[![GoDoc](https://godoc.org/github.com/abhiyerra/sarpa/client?status.svg)](https://godoc.org/github.com/abhiyerra/sarpa/client)

# Sarpa

Sarpa is a HTTP/HTTPS proxy for web services which use service
discovery tools like etcd. Right now it supports nginx and etcd.


# Configuration

The configuration is JSON definition of key values:


    {
        "etcd_hosts": ["http://127.0.0.1:4001"],
        "restart_cmd": "service nginx restart"
        "services": [
            {
                "service_name": "treemap",
                "hosts": ["treemap.org", "www.treemap.org"]
            },
            {
                "service_name": "forestly",
                "hosts": ["forestly.org", "www.forestly.org"]
            }
        ]
    }


etcd_hosts
: An array of etcd hosts to connect to.
restart_cmd
: The command to restart nginx
services
: A list of objects defining the service_name and an array of hostnames.

# Client

There is a Sarpa client in the client directory

    go get -u github.com/abhiyerra/sarpa/client

To run the code in you Go application call:

  SarpaUpdater()

Please look at the
[godocs](http://godoc.org/github.com/abhiyerra/sarpa/client) for the
client to understand how it works.
