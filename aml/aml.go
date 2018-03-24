package aml

import (
	"github.com/windler/awesm/aml/instructions"
)

//Parser parses a *.aml file
type Parser interface {
	AddInstructionFactory(keyword string, factory InstructionFactory)
	AddForkJoinFactory(keyword string, factory ForkJoinFactory)
	Parse() (File, error)
}

//InstructionFactory creates instructions based on patterns
type InstructionFactory interface {
	New(name string) *instructions.AMLInstruction
	GetPattern() string
}

//ForkJoinFactory creates fork join nodes
type ForkJoinFactory interface {
	NewForkNode(name string) *instructions.AMLInstruction
	NewJoinNode(name string, forkNode *instructions.AMLInstruction) *instructions.AMLInstruction
	GetPattern() string
	ProvidesPathLabels() bool
}
