package entities

type Cluster struct {
	Leader      bool `json:"leader"`
	AliveRemote bool `json:"alive"`
}
