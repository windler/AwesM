package instructions

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

type IfInstructionFactory struct{}

func (f IfInstructionFactory) New(name string) *AMLInstruction {
	return &AMLInstruction{
		Name: getNodeName(name),
		EdgeOptions: map[string]string{
			"label": getLabelName(name),
		},
	}
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
		Predecessors: getJoinNodePredecessors(name, forkNode),
	}
}

func getLabelName(name string) string {
	r := regexp.MustCompile("\\?{1,2}\\[(.+)\\].+")
	return " [" + r.FindStringSubmatch(name)[1] + "]"
}

func getNodeName(name string) string {
	r := regexp.MustCompile("\\?{1,2}\\[.+\\](.+)")
	return r.FindStringSubmatch(name)[1]
}

func getJoinNodePredecessors(name string, ins *AMLInstruction) []*AMLInstruction {
	if strings.Contains(name, "??") {
		return []*AMLInstruction{ins}
	}
	return []*AMLInstruction{}
}

func (f IfInstructionFactory) GetPattern() string {
	return "\\?{1,2}\\[.+\\]"
}
