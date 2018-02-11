package instructions

import "github.com/windler/dotgraph/graph"

type NoForkFactory struct{}

func (f *NoForkFactory) NewJoinNode(name string, forkNode *AMLInstruction) *AMLInstruction {
	return nil
}

func (f *NoForkFactory) NewForkNode(name string) *AMLInstruction {
	return nil
}

type AMLInstruction struct {
	Name         string
	Predecessors []*AMLInstruction
	NodeOptions  map[string]string
	EdgeOptions  map[string]string
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
