package instructions

import "github.com/windler/dotgraph/graph"

type AMLInstruction struct {
	Name          string
	Predecesors   []*AMLInstruction
	NodeOptions   map[string]string
	EdgeOptions   map[string]string
	PathBeginning bool
	PathJoinNode  *AMLInstruction
}

func (i *AMLInstruction) GetPredecessors() []*AMLInstruction {
	return i.Predecesors
}

func (i *AMLInstruction) GetNodeOptions() graph.DotGraphOptions {
	return i.NodeOptions
}

func (i *AMLInstruction) GetEdgeOptions() graph.DotGraphOptions {
	return i.EdgeOptions
}

func (i *AMLInstruction) GetNodeName() string {
	return i.Name
}

func (i *AMLInstruction) IsPathBeginning() bool {
	return i.PathBeginning
}

func (i *AMLInstruction) GetPathJoinNode() *AMLInstruction {
	return i.PathJoinNode
}
