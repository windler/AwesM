package instructions

type EndInstructionFactory struct{}

func (f EndInstructionFactory) New(name string, predecessor, parent *AMLInstruction) AMLInstruction {
	return AMLInstruction{
		Name:        name,
		Predecesors: []*AMLInstruction{predecessor},
		NodeOptions: map[string]string{
			"shape":     "doublecircle",
			"label":     "",
			"style":     "filled",
			"fillcolor": "#111111",
			"height":    "0.3",
		},
		EdgeOptions: make(map[string]string),
	}
}

func (f EndInstructionFactory) GetPattern() string {
	return "#end"
}
