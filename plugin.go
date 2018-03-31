package vexlink

// Plugin represents a Vexlink plugin that provides nodes.
type Plugin struct {
	// Name is a unique identifier for this plugin. It should consist only of lowercase letters and numbers.
	Name string
	// Version is the build version of the plugin.
	Version string
	// Nodes is a map of node names to their prototype Node objects.
	Nodes map[string]Node
}
