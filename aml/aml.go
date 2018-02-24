package aml

import (
	"github.com/windler/awesm/aml/instructions"
)

//Parser parses a *.aml file
type Parser interface {
	AddFactory(keyword string, factory InstructionFactory)
	Parse() (AMLFile, error)
}

//InstructionFactory creates instructions based on patterns
type InstructionFactory interface {
	New(name string) *instructions.AMLInstruction
	NewForkNode(name string) *instructions.AMLInstruction
	NewJoinNode(name string, forkNode *instructions.AMLInstruction) *instructions.AMLInstruction
	GetPattern() string
}
