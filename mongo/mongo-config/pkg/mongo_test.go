package mongo

import (
	"testing"
)

func TestYamlParse(t *testing.T) {
	yamlString := `
replicaSets:
  - replicaSet:
    name : rs0
    members :
      - arbiterOnly: "false"
        host: "host1"
      - arbiterOnly: "false"
        host: "host2"
      - arbiterOnly: "true"
        host: "host3"
  - replicaSet:
    name : rs1
    members : 
      - arbiterOnly: "false"
        host: "host1"
      - arbiterOnly: "true"
        host: "host2"
      - arbiterOnly: "true"
        host: "host3"
`

	cluster, err := parseString(yamlString)
	if err != nil {
		t.Fatalf("Got error but shouldn't have: %+v", err)
	}
	t.Log("cluster = %+v", cluster)

	bytes := []byte(yamlString)
	cluster, err = parseBytes(bytes)

	if err != nil {
		t.Fatalf("Got error but shouldn't have: %+v", err)
	}
	t.Log("cluster = %+v", cluster)

	if len(cluster.ReplicaSets) != 2 {
		t.Fatalf("Number of replica sets should be 1")
	}

	replicaSetInfo := cluster.ReplicaSets[0]
	if replicaSetInfo.Name != "rs0" {
		t.Fatalf("replicaSetInfo.Name should be 'rs0' but is '%s'", replicaSetInfo.Name)
	}
}

func TestBuildReplSetBson(t *testing.T) {
	yamlString := `
replicaSets:
  - replicaSet:
    name : rs0
    members :
      - arbiterOnly: "false"
        host: "host1"
      - arbiterOnly: "false"
        host: "host2"
      - arbiterOnly: "true"
        host: "host3"
  - replicaSet:
    name : rs1
    members : 
      - arbiterOnly: "false"
        host: "host1"
      - arbiterOnly: "true"
        host: "host2"
      - arbiterOnly: "true"
 `
	cluster, err := parseString(yamlString)
	if err != nil {
		t.Fatalf("Got error but shouldn't have: %+v", err)
	}

	doc, primaryHost := buildReplSetInitiateBson(cluster.ReplicaSets[0])
	if primaryHost != "host1" {
		t.Fatalf("primaryHost should == 'host1', but is '%s'", primaryHost)
	}

	t.Logf("primaryHost = '%s'", primaryHost)
	t.Logf("doc = '%+v'", doc)
}

func TestdReplSetInit(t *testing.T) {
	yamlString := `
replicaSets:
  - replicaSet:
    name : rs0
    members :
      - arbiterOnly: "false"
        host: "localhost"
      - arbiterOnly: "false"
        host: "host2"
      - arbiterOnly: "true"
        host: "host3"
  - replicaSet:
    name : rs1
    members : 
      - arbiterOnly: "false"
        host: "host1"
      - arbiterOnly: "true"
        host: "host2"
      - arbiterOnly: "true"
 `
	cluster, err := parseString(yamlString)
	if err != nil {
		t.Fatalf("Got error but shouldn't have: %+v", err)
	}

	doc, primaryHost := buildReplSetInitiateBson(cluster.ReplicaSets[0])
	if primaryHost != "host1" {
		t.Fatalf("primaryHost should == 'host1', but is '%s'", primaryHost)
	}

	t.Logf("primaryHost = '%s'", primaryHost)
	t.Logf("doc = '%+v'", doc)
}

func TestInvalidYamlParse(t *testing.T) {
	yamlString := `
replicaSets:
  - replicaSet:
    name : rs0
    members :
      - arbiterOnly: "false"
        host: "host1"
      - arbiterOnly: "false"
        host: "host2"
      - arbiterOnly: "true"
        host: "host3"
  - replicaSet:
    name : rs1
    members : 
      - arbiterOnly: "false"
        host: "host1"
      - arbiterOnly: "true"
        host: "host2"
      - arbiterOnly: "true"
 `

	cluster, err := parseString(yamlString)
	if err != nil {
		t.Fatalf("Got error but shouldn't have: %+v", err)
	}
	t.Log("cluster = %+v", cluster)

	bytes := []byte(yamlString)
	cluster, err = parseBytes(bytes)

	if err != nil {
		t.Fatalf("Got error but shouldn't have: %+v", err)
	}
	t.Log("cluster = %+v", cluster)

	if len(cluster.ReplicaSets) != 2 {
		t.Fatalf("Number of replica sets should be 1")
	}

	replicaSetInfo := cluster.ReplicaSets[0]
	if replicaSetInfo.Name != "rs0" {
		t.Fatalf("replicaSetInfo.Name should be 'rs0' but is '%s'", replicaSetInfo.Name)
	}
}

func TestFileParse(t *testing.T) {
	shardInfo, err := ReadYamlFile("test_file.yaml")
	if err != nil {
		t.Fatalf("Got error but shouldn't have: %+v", err)
	}
	t.Log("shardInfo = %+v", shardInfo)
}

func TestToYaml(t *testing.T) {

}
