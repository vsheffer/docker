package shard

import (
	"flag"
	"github.com/golang/glog"
	"gopkg.in/mgo.v2"
)

type ShardInfo struct {
	ReplicaSets []ReplicaHostInfo `yaml: replicaSets`
}

type ReplicaHostInfo struct {
	Hostname       string `yaml: hostname`
	ReplicaSetName string `yaml: replicaSetName`
}
