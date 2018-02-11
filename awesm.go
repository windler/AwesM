package main

import (
	"fmt"

	"github.com/windler/dotgraph/renderer"

	"github.com/windler/awesm/aml"
	"github.com/windler/awesm/aml/instructions"
)

func main() {
	amlParser := aml.NewFileParser("examples/simple_test.aml")
	amlParser.AddFactory(instructions.StartInstructionFactory{})
	amlParser.AddFactory(instructions.EndInstructionFactory{})
	amlParser.AddFactory(instructions.IfInstructionFactory{})
	amlParser.AddFactory(instructions.ParallelInstructionFactory{})
	amlParser.AddFactory(instructions.ActivityInstructionFactory{})

	aml, err := amlParser.Parse()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	r := renderer.PNGRenderer{
		HomeDir:    "",
		OutputFile: "examples/simple_test.png",
		Prefix:     "",
	}
	r.Render(aml.CreateDotGraph().String())
}
