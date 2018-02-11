package aml

import (
	"github.com/windler/awesm/aml/instructions"
)

type AMLParser interface {
	AddFactory(keyword string, factory AMLInstructionFactory)
	Parse() (AMLFile, error)
}

type AMLInstructionFactory interface {
	New(name string) *instructions.AMLInstruction
	NewForkNode(name string) *instructions.AMLInstruction
	NewJoinNode(name string, forkNode *instructions.AMLInstruction) *instructions.AMLInstruction
	GetPattern() string
}
