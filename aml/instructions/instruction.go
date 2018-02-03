package instructions

import "github.com/windler/godepg/dotgraph"

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

func (i *AMLInstruction) GetNodeOptions() dotgraph.DotGraphOptions {
	return i.NodeOptions
}

func (i *AMLInstruction) GetEdgeOptions() dotgraph.DotGraphOptions {
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
