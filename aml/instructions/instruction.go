package instructions

import "github.com/windler/dotgraph/graph"

//NoForkFactory is a factory that does not provide fork- and join-nodes
type NoForkFactory struct{}

//NewJoinNode create a join node which is added a the end of the path
func (f *NoForkFactory) NewJoinNode(name string, forkNode *AMLInstruction) *AMLInstruction {
	return nil
}

//NewForkNode create a fork node which is added a the beginnen of the path
func (f *NoForkFactory) NewForkNode(name string) *AMLInstruction {
	return nil
}

//AMLInstruction is a instruction within an *.aml file
type AMLInstruction struct {
	Name         string
	Predecessors []*AMLInstruction
	NodeOptions  map[string]string
	EdgeOptions  map[string]string
}

//GetPredecessors gets the predecessor nodes
func (i *AMLInstruction) GetPredecessors() []*AMLInstruction {
	return i.Predecessors
}

//GetNodeOptions gets the style attributes for this node
func (i *AMLInstruction) GetNodeOptions() graph.DotGraphOptions {
	return i.NodeOptions
}

//GetEdgeOptions gets the style attributes for an edge from this node
func (i *AMLInstruction) GetEdgeOptions() graph.DotGraphOptions {
	return i.EdgeOptions
}

//GetNodeName gets the name of this node
func (i *AMLInstruction) GetNodeName() string {
	return i.Name
}
