package instructions

import (
	"strings"
)

//StartInstructionFactory creates a starting-node
type StartInstructionFactory struct {
	*NoForkFactory
}

//New creates a new instrcution
func (f StartInstructionFactory) New(name string) *AMLInstruction {
	return &AMLInstruction{
		Name:         strings.Replace(name, "?", "cond_", -1),
		Predecessors: []*AMLInstruction{},
		NodeOptions: map[string]string{
			"shape":     "circle",
			"label":     "",
			"style":     "filled",
			"fillcolor": "#111111",
			"height":    "0.3",
		},
		EdgeOptions: make(map[string]string),
	}
}

//GetPattern get the pattern this factory can handle
func (f StartInstructionFactory) GetPattern() string {
	return "#start"
}
