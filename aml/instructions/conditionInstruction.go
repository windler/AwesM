package instructions

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

type IfInstructionFactory struct{}

func (f IfInstructionFactory) New(name string) *AMLInstruction {
	activity := ActivityInstructionFactory{}.New(getCondNodeName(name))
	activity.EdgeOptions["label"] = getCondLabelName(name)
	return activity
}

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

func (f IfInstructionFactory) GetPattern() string {
	return "\\?{1,2}\\[.+\\]"
}
