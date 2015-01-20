package main

import (
	"flag"
	mc "github.com/vsheffer/docker/mongo/mongo-config/pkg"
	"log"
)

func main() {
	yamlFile := flag.String("yamlFile", "", "")
	cluster, err := mc.NewCluster(*yamlFile)

	if err != nil {
		log.Fatalf("err %+v", err)
	}
	cluster.ReplSetInitiate()
}
