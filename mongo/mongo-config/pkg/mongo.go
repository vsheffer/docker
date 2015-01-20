// Execute all Mongo commands to both initiate replica sets and then turn on sharding for a Mongo cluster
// deployed using Kubernetes.
//
// Configuration information about the cluster is read in from a YAML file.
//
// This executable gets added to a simple Docker container that is suitable to run in a Kubernetes pod.
package mongo

import (
	"github.com/ghodss/yaml"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/validator.v2"
	"io/ioutil"
	"log"
	"time"
)

type Cluster struct {
	ReplicaSets []ReplicaSetInfo `json:"replicaSets"`
}

type ReplicaSetInfo struct {
	Name    string             `json:"name"  validate:"nonzero"`
	Members []ReplicaSetMember `json:"members"`
}

type Boolean struct {
	value string
}

type ReplicaSetMember struct {
	ArbiterOnly string `json:"arbiterOnly"`
	Host        string `json:"host"  validate:"nonzero"`
}

func isTrue(s string) bool {
	if len(s) == 0 {
		return false
	}
	return s == "true"
}

func ReadYamlFile(filename string) (*Cluster, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return parseBytes(bytes)
}

func parseString(yamlString string) (*Cluster, error) {
	return parseBytes([]byte(yamlString))
}

func parseBytes(bytes []byte) (*Cluster, error) {
	shardInfo := &Cluster{}
	err := yaml.Unmarshal(bytes, shardInfo)
	return shardInfo, err
}

func NewCluster(shardConfigFile string) (*Cluster, error) {
	Cluster, _ := ReadYamlFile(shardConfigFile)
	err := validator.Validate(Cluster)
	if err != nil {
		return nil, err
	}

	return Cluster, nil
}

func buildReplSetInitiateBson(replicaSetInfo ReplicaSetInfo) (bson.D, string) {
	members := make([]bson.D, 0)
	var primaryHost string
	for memberId, replicaSetMember := range replicaSetInfo.Members {
		doc := bson.D{{"_id", memberId}, {"host", replicaSetMember.Host}, {"arbiterOny", isTrue(replicaSetMember.ArbiterOnly)}}
		members = append(members, doc)
		if memberId == 0 {
			primaryHost = replicaSetMember.Host
		}
	}

	return bson.D{{"replSetInitiate", bson.D{{"_id", replicaSetInfo.Name}, {"members", members}}}}, primaryHost
}

func (c *Cluster) ReplSetInitiate() error {
	var session *mgo.Session
	var err error

	for _, replSet := range c.ReplicaSets {
		doc, primaryHostname := buildReplSetInitiateBson(replSet)
		log.Printf("doc = %+v", doc)

		var result interface{}
		log.Printf("Attempting connection to %s", primaryHostname)
		session, err = mgo.DialWithTimeout(primaryHostname, time.Duration(10)*time.Second)
		if err != nil {
			return err
		}

		err := session.Run(doc, &result)
		log.Printf("err = %+v, result = %+v", err, result)
		if err != nil {
			return err
		}
	}
	return nil
}
