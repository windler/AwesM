package instructions

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

type IfInstructionFactory struct{}

func (f IfInstructionFactory) New(name string) AMLInstruction {
	randName := strconv.FormatInt(int64(rand.Int()), 10)

	ins := AMLInstruction{
		Name: getNodeName(name),
		EdgeOptions: map[string]string{
			"label": getLabelName(name),
		},
	}

	forkNode := &AMLInstruction{
		Name: "cond_" + randName,
		NodeOptions: map[string]string{
			"shape":     "diamond",
			"label":     "",
			"fillcolor": "#111111",
		},
		EdgeOptions: make(map[string]string),
	}

	joinNode := &AMLInstruction{
		Name: "join_" + randName,
		NodeOptions: map[string]string{
			"shape":     "diamond",
			"label":     "",
			"fillcolor": "#111111",
		},
		EdgeOptions:  make(map[string]string),
		Predecessors: getJoinNodePredecessors(name, forkNode),
	}

	forkNode.PathJoinNode = joinNode
	ins.PathForkNode = forkNode

	return ins
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
