package instructions

import (
	"math/rand"
	"strconv"
	"strings"
)

type IfInstructionFactory struct{}

func (f IfInstructionFactory) New(name string, predecessor, parent *AMLInstruction) AMLInstruction {
	randName := strconv.FormatInt(int64(rand.Int()), 10)
	ins := AMLInstruction{
		Name:        "cond_" + randName,
		Predecesors: []*AMLInstruction{predecessor},
		NodeOptions: map[string]string{
			"shape":     "diamond",
			"label":     "",
			"fillcolor": "#111111",
		},
		EdgeOptions: make(map[string]string),
	}

	pathJoinNodePredecessors := getJoinNodePredecessors(name, &ins)

	ins.PathJoinNode = &AMLInstruction{
		Name: "join_" + randName,
		NodeOptions: map[string]string{
			"shape":     "diamond",
			"label":     "",
			"fillcolor": "#111111",
		},
		EdgeOptions: make(map[string]string),
		Predecesors: pathJoinNodePredecessors,
	}

	return ins
}

func getJoinNodePredecessors(name string, ins *AMLInstruction) []*AMLInstruction {
	if strings.Contains(name, "??") {
		return []*AMLInstruction{ins}
	}
	return []*AMLInstruction{}
}

func (f IfInstructionFactory) GetPattern() string {
	return "\\?"
}
