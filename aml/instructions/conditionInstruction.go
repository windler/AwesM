package instructions

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

//IfInstructionFactory creates condition nodes
type IfInstructionFactory struct{}

//New creates a new instrcution
func (f IfInstructionFactory) New(name string) *AMLInstruction {
	activity := ActivityInstructionFactory{}.New(getCondNodeName(name))
	activity.EdgeOptions["label"] = getCondLabelName(name)
	return activity
}

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
		EdgeOptions: make(map[string]string),
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

func getCondLabelName(name string) string {
	r := regexp.MustCompile("\\?{1,2}\\[(.+)\\].+")
	return " [" + r.FindStringSubmatch(name)[1] + "]"
}

func getCondNodeName(name string) string {
	r := regexp.MustCompile("\\?{1,2}\\[.+\\](.+)")
	return r.FindStringSubmatch(name)[1]
}

func getCondJoinNodePredecessors(name string, ins *AMLInstruction) []*AMLInstruction {
	if strings.Contains(name, "??") {
		return []*AMLInstruction{ins}
	}
	return []*AMLInstruction{}
}

//GetPattern get the pattern this factory can handle
func (f IfInstructionFactory) GetPattern() string {
	return "\\?{1,2}\\[.+\\]"
}
