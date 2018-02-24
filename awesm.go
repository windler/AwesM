package main

import (
	"fmt"
	"os"

	"github.com/windler/dotgraph/renderer"

	"github.com/windler/awesm/aml"
	"github.com/windler/awesm/aml/instructions"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) != 1 {
		fmt.Println("no file provided.")
		return
	}

	amlParser := aml.NewFileParser(argsWithoutProg[0])
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
		OutputFile: argsWithoutProg[0] + ".png",
		Prefix:     "",
	}
	r.Render(aml.CreateDotGraph().String())
}
