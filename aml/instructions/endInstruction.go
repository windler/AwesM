package instructions

type EndInstructionFactory struct{}

func (f EndInstructionFactory) New(name string) AMLInstruction {
	return AMLInstruction{
		Name: name,
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
