package instructions

import (
	"math/rand"
	"strconv"
)

//IfInstructionFactory creates condition nodes
type IfInstructionFactory struct{}

//NewForkNode create a fork node which is added a the beginnen of the path
func (f IfInstructionFactory) NewForkNode(name string) *AMLInstruction {
	randName := strconv.FormatInt(int64(rand.Int()), 10)
	return &AMLInstruction{
		Name: "cond_" + randName,
		NodeOptions: map[string]string{
			"shape":     "diamond",
			"label":     "",
			"fillcolor": "#111111",
		},
		EdgeOptions:  make(map[string]string),
		Predecessors: []*AMLInstruction{},
	}
}

//NewJoinNode create a join node which is added a the end of the path
func (f IfInstructionFactory) NewJoinNode(name string, forkNode *AMLInstruction) *AMLInstruction {
	randName := strconv.FormatInt(int64(rand.Int()), 10)
	return &AMLInstruction{
		Name: "join_" + randName,
		NodeOptions: map[string]string{
			"shape":     "diamond",
			"label":     "",
			"fillcolor": "#111111",
		},
		EdgeOptions:  make(map[string]string),
		Predecessors: getCondJoinNodePredecessors(name, forkNode),
	}
}

func getCondJoinNodePredecessors(name string, ins *AMLInstruction) []*AMLInstruction {
	if name == "ifopt" {
		return []*AMLInstruction{ins}
	}
	return []*AMLInstruction{}
}

//GetPattern get the pattern this factory can handle
func (f IfInstructionFactory) GetPattern() string {
	return "if|ifopt"
}

//ProvidesPathLabels determines wether this nodes has sub paths
func (f IfInstructionFactory) ProvidesPathLabels() bool {
	return true
}
