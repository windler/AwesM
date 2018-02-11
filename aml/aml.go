package aml

import (
	"github.com/windler/awesm/aml/instructions"
)

type AMLParser interface {
	AddFactory(keyword string, factory AMLInstructionFactory)
	Parse() (AMLFile, error)
}

type AMLInstructionFactory interface {
	New(name string) instructions.AMLInstruction
	GetPattern() string
}
