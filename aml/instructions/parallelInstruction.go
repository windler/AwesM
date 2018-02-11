package instructions

import (
	"math/rand"
	"regexp"
	"strconv"
)

type ParallelInstructionFactory struct{}

func (f ParallelInstructionFactory) New(name string) *AMLInstruction {
	return ActivityInstructionFactory{}.New(getParallelNodeName(name))
}

func (f ParallelInstructionFactory) NewForkNode(name string) *AMLInstruction {
	randName := strconv.FormatInt(int64(rand.Int()), 10)
	return &AMLInstruction{
		Name: "cond_" + randName,
		NodeOptions: map[string]string{
			"shape":     "rectangle",
			"label":     "",
			"fillcolor": "#111111",
			"height":    "0.1",
			"width":     "2",
		},
		EdgeOptions: make(map[string]string),
	}
}

func (f ParallelInstructionFactory) NewJoinNode(name string, forkNode *AMLInstruction) *AMLInstruction {
	randName := strconv.FormatInt(int64(rand.Int()), 10)
	return &AMLInstruction{
		Name: "join_" + randName,
		NodeOptions: map[string]string{
			"shape":     "rectangle",
			"label":     "",
			"fillcolor": "#111111",
			"height":    "0.1",
			"width":     "2",
		},
		EdgeOptions: make(map[string]string),
	}
}

func getParallelNodeName(name string) string {
	r := regexp.MustCompile("\\|(.+)")
	return r.FindStringSubmatch(name)[1]
}

func (f ParallelInstructionFactory) GetPattern() string {
	return "\\|.+"
}
