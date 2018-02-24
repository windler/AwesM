package instructions

import (
	"math/rand"
	"regexp"
	"strconv"
)

//ParallelInstructionFactory create a parallel-node
type ParallelInstructionFactory struct{}

//New creates a new instrcution
func (f ParallelInstructionFactory) New(name string) *AMLInstruction {
	return ActivityInstructionFactory{}.New(getParallelNodeName(name))
}

//NewForkNode create a fork node which is added a the beginnen of the path
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

//NewJoinNode create a join node which is added a the end of the path
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

//GetPattern get the pattern this factory can handle
func (f ParallelInstructionFactory) GetPattern() string {
	return "\\|.+"
}
