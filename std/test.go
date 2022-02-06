package main

import (
	"fmt"
)

func main() {

	//r := fmt.Sprintf("https://kubernetes.default.svc.%s:6443", "cluster.local")
	//fmt.Println(r)

    config := Config {
		Kind: "Config",
		Clusters: map[string]*Cluster{
			"A":{
				Server : fmt.Sprintf("https://kubernetes.default.svc.%s:6443", "cluster.local"),
				Data: "test",
			},
		},
	}

	fmt.Println(config.Clusters["A"].Server,config.Clusters["A"].Data)

}

type Cluster struct {
	Server string
	Data   string
}

type Config struct {
	Kind     string
	Clusters map[string]*Cluster
}
