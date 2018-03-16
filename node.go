package vexlink

type NodeInformation struct {
	// Options maps option names to their default values.
	Options map[string]interface{}

	// Inputs is a list of inputs this node accepts.
	Inputs []string

	// Outputs is a list of outputs this node returns.
	Outputs []string
}

type Node interface {
	// GetInformation returns information about the node. This should always return the same value.
	GetInformation() NodeInformation

	// Run runs the node. The node processes input and returns output. The node should terminate its loop as soon as data is received on the stop channel.
	Run(options map[string]interface{}, input chan map[string]interface{}, output chan map[string]interface{}, stop chan struct{})
}
