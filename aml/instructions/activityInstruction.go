package instructions

type ActivityInstructionFactory struct{}

func (f ActivityInstructionFactory) New(name string, predecessor, parent *AMLInstruction) AMLInstruction {
	return AMLInstruction{
		Name:        name,
		Predecesors: []*AMLInstruction{predecessor},
		NodeOptions: make(map[string]string),
		EdgeOptions: make(map[string]string),
	}
}

func (f ActivityInstructionFactory) GetPattern() string {
	return ".*"
}
