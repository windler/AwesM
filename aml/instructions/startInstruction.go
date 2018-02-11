package instructions

import (
	"strings"
)

type StartInstructionFactory struct{}

func (f StartInstructionFactory) New(name string) AMLInstruction {
	return AMLInstruction{
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

func (f StartInstructionFactory) GetPattern() string {
	return "#start"
}
