package instructions

import (
	"math/rand"
	"regexp"
	"strconv"

	"github.com/windler/awesm/aml/constants"
)

//ParallelInstructionFactory create a parallel-node
type ParallelInstructionFactory struct {
	Orientation string
}

//NewForkNode create a fork node which is added a the beginnen of the path
func (f ParallelInstructionFactory) NewForkNode(name string) *AMLInstruction {
	randName := strconv.FormatInt(int64(rand.Int()), 10)

	height := "0.1"
	width := "2"
	if f.Orientation == constants.LeftRight {
		width = "0.1"
		height = "2"
	}

	return &AMLInstruction{
		Name: "cond_" + randName,
		NodeOptions: map[string]string{
			"shape":     "rectangle",
			"label":     "",
			"fillcolor": "#111111",
			"height":    height,
			"width":     width,
		},
		EdgeOptions:  make(map[string]string),
		Predecessors: []*AMLInstruction{},
	}
}

//NewJoinNode create a join node which is added a the end of the path
func (f ParallelInstructionFactory) NewJoinNode(name string, forkNode *AMLInstruction) *AMLInstruction {
	randName := strconv.FormatInt(int64(rand.Int()), 10)

	height := "0.1"
	width := "2"
	if f.Orientation == constants.LeftRight {
		width = "0.1"
		height = "2"
	}

	return &AMLInstruction{
		Name: "join_" + randName,
		NodeOptions: map[string]string{
			"shape":     "rectangle",
			"label":     "",
			"fillcolor": "#111111",
			"height":    height,
			"width":     width,
		},
		EdgeOptions:  make(map[string]string),
		Predecessors: []*AMLInstruction{},
	}
}

func getParallelNodeName(name string) string {
	r := regexp.MustCompile("\\|(.+)")
	return r.FindStringSubmatch(name)[1]
}

//GetPattern get the pattern this factory can handle
func (f ParallelInstructionFactory) GetPattern() string {
	return "fork"
}

//ProvidesPathLabels determines wether this nodes has sub paths
func (f ParallelInstructionFactory) ProvidesPathLabels() bool {
	return false
}
