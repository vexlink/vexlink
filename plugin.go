package vexlink

type Plugin struct {
	Name    string
	Version string
	Nodes   map[string]Node
}
