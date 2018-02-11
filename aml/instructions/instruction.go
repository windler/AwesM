package instructions

import "github.com/windler/dotgraph/graph"

type AMLInstruction struct {
	Name         string
	Predecessors []*AMLInstruction
	NodeOptions  map[string]string
	EdgeOptions  map[string]string
	PathForkNode *AMLInstruction
	PathJoinNode *AMLInstruction
}

func (i *AMLInstruction) GetPredecessors() []*AMLInstruction {
	return i.Predecessors
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
	return i.PathForkNode != nil
}

func (i *AMLInstruction) GetPathJoinNode() *AMLInstruction {
	return i.PathJoinNode
}

func (i *AMLInstruction) GetPathForkNode() *AMLInstruction {
	return i.PathForkNode
}
