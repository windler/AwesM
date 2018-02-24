package instructions

//EndInstructionFactory creates end-nodes
type EndInstructionFactory struct {
	*NoForkFactory
}

//New creates a new instrcution
func (f EndInstructionFactory) New(name string) *AMLInstruction {
	return &AMLInstruction{
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

//GetPattern get the pattern this factory can handle
func (f EndInstructionFactory) GetPattern() string {
	return "#end"
}
