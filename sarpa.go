package main

import (
	"encoding/json"
	"fmt"
	"github.com/coreos/go-etcd/etcd"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/s3"
	"log"
	"os"
	"strings"
)

const (
	EtcdSarpaPath = "/sarpa"
)

var (
	config *Config
)

type Config struct {
	etcdClient *etcd.Client `json:"-"`
	awsAuth    aws.Auth     `json:"-"`
	s3Auth     *s3.S3       `json:"-"`

	EtcdHosts []string            `json:"etcd_hosts"`
	S3Bucket  string              `json:"-"`
	Services  map[string][]string `json:"services"`
}

func (c *Config) GetServices() {
	newServices := make(map[string][]string)

	// Get the nodes and their values.
	resp, err := c.etcdClient.Get(EtcdSarpaPath, true, true)
	if err != nil {
		log.Println(err)
		c.Services = newServices
		return
	}

	var newProxies []string
	for _, n := range resp.Node.Nodes {
		log.Printf("%s: %s\n", n.Key, n.Value)

		keySplit := strings.SplitAfter(n.Key, "/")
		serviceName := keySplit[len(keySplit)-1]

		if len(n.Nodes) > 0 {
			for _, k := range n.Nodes {
				log.Printf("%s: %s\n", k.Key, k.Value)
				newServices[serviceName] = append(newServices[serviceName], k.Value)
			}
		}

		newProxies = append(newProxies, n.Value)
	}

	c.Services = newServices
}

func (c *Config) jsonedServices() []byte {
	data, err := json.Marshal(c.Services)
	if err != nil {
		log.Println(err)
	}

	return data
}

func (c *Config) UploadToS3() {
	if c.s3Auth == nil {
		c.s3Auth = s3.New(c.awsAuth, aws.USEast)
	}
	bucket := c.s3Auth.Bucket(c.S3Bucket)

	err := bucket.Put("discovery.js", []byte(fmt.Sprintf("var SarpaServiceDiscovery = %s;", c.jsonedServices())), "text/javascript", s3.PublicRead, s3.Options{})
	if err != nil {
		log.Println(err)
	}
}

func (c *Config) EtcdConnect(etcdHosts []string) {
	c.EtcdHosts = etcdHosts
	c.etcdClient = etcd.NewClient(c.EtcdHosts)
}

func (c *Config) AwsConnect() {
	auth, err := aws.EnvAuth()
	if err != nil {
		log.Println(err)
	}

	c.awsAuth = auth
}

func (c *Config) StartWatchmen() {
	log.Println("Starting watchman.")

	for {
		log.Println("Waiting for an update...")

		// Watch for changes to the values
		watchChan := make(chan *etcd.Response)
		go c.etcdClient.Watch(EtcdSarpaPath, 0, true, watchChan, nil)

		select {
		case r := <-watchChan:
			log.Printf("updated keys: %s: %s\n", r.Node.Key, r.Node.Value)
			c.GetServices()

			log.Println("Uploading to S3.")
			c.UploadToS3()
		}
	}
}

func init() {
	etcdHost := os.Getenv("ETCD_HOSTS")
	if etcdHost == "" {
		etcdHost = "http://127.0.0.1:4001"
	}

	config = &Config{}
	config.EtcdConnect([]string{etcdHost})
	config.AwsConnect()

	config.S3Bucket = os.Getenv("SARPA_BUCKET")
	log.Println("S3 Bucket", config.S3Bucket)
}

func main() {
	// Need to run it the first time to reset the stale config.
	config.GetServices()
	config.UploadToS3()

	config.StartWatchmen()
}
