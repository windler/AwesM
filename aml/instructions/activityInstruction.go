package instructions

//ActivityInstructionFactory creates Activity-Nodes
type ActivityInstructionFactory struct{}

//New creates a new instrcution
func (f ActivityInstructionFactory) New(name string) *AMLInstruction {
	return &AMLInstruction{
		Name:         name,
		NodeOptions:  make(map[string]string),
		EdgeOptions:  make(map[string]string),
		Predecessors: []*AMLInstruction{},
	}
}

//GetPattern get the pattern this factory can handle
func (f ActivityInstructionFactory) GetPattern() string {
	return ".*"
}
