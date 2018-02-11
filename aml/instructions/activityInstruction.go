package instructions

type ActivityInstructionFactory struct {
	*NoForkFactory
}

func (f ActivityInstructionFactory) New(name string) *AMLInstruction {
	return &AMLInstruction{
		Name:        name,
		NodeOptions: make(map[string]string),
		EdgeOptions: make(map[string]string),
	}
}

func (f ActivityInstructionFactory) GetPattern() string {
	return ".*"
}
