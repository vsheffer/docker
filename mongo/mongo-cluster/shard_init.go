// Execute all Mongo commands to both initiate replica sets and then turn on sharding for a Mongo cluster
// deployed using Kubernetes.
//
// Configuration information about the cluster is read in from a YAML file.
//
// This executable gets added to a simple Docker container that is suitable to run in a Kubernetes pod.
package main

import (
	"flag"
	"github.com/vsheffer/docker/mongo/pkg/mongocluster"
	"log"
)

func main() {
	mongoUrl := flag.String("mongoUrl", "", "The URL for the mongo.")
	filename := flag.String("f", "shard_config.yaml", "The filename of the shard config file.")
	flag.Parse()
	shard, err := mongocluster.NewShard(*mongoUrl, *filename)
	if err != nil {
		log.Fatalf("ReplSetInitiate failed: %+v", err)
	}

	defer shard.Session.Close()

	err = shard.ReplSetInitiate()
	if err != nil {
		log.Fatalf("ReplSetInitiate failed: %+v", err)
	}
}
