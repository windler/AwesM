package instructions

import (
	"strings"
)

type StartInstructionFactory struct{}

func (f StartInstructionFactory) New(name string, predecessor, parent *AMLInstruction) AMLInstruction {
	return AMLInstruction{
		Name:        strings.Replace(name, "?", "cond_", -1),
		Predecesors: []*AMLInstruction{},
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
