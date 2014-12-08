[![GoDoc](https://godoc.org/github.com/abhiyerra/sarpa/client?status.svg)](https://godoc.org/github.com/abhiyerra/sarpa/client)

# Sarpa

Sarpa is a poor man's service discovery system. It is made to be used
with Docker instances. So when an Docker instance is started up it
reports itself on etcd with the public ip and port to connect
to. Sarpa watches for these key changes. And updates a file in S3
which has key values of service to machines.

# Running

    export AWS_ACCESS_KEY_ID=<your access key id>
    export AWS_SECRET_ACCESS_KEY=<your secret key>
    export ETCD_HOSTS="http://127.0.0.1:4001"
    export SARPA_BUCKET="discovery.forestly.org"
    ./sarpa

# Etcd Keys

The keys for the discovery looks like the following. It always looks
in the /sarpa directory. The next key is the s3 bucket to push the
configuration file to. The value after that is the service short name.

In that directory you can create children which include the public_host:port
so the clients can connect to them.

    /sarpa/:service/:machine_id

Example:

    etcdctl set /sarpa/treely/1 "104.236.59.43:3001"

Example with fleet for the service.


    [Unit]
    Description=treely

    [Service]
    ExecStartPre=-/usr/bin/docker kill treely-%i
    ExecStartPre=-/usr/bin/docker rm treely-%i
    ExecStart=/usr/bin/docker run --rm --name treely-%i -p 8080:8080 forestly/treely
    ExecStartPost=/usr/bin/etcdctl set /sarpa/treely/%m %H:%i
    ExecStop=/usr/bin/docker stop treely-%i
    ExecStopPost=/usr/bin/etcdctl rm /sarpa/treemly/%m

    [X-Fleet]
    Conflicts=treely@*.service
