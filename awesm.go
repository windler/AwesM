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
	amlParser.AddInstructionFactory(instructions.StartInstructionFactory{})
	amlParser.AddInstructionFactory(instructions.EndInstructionFactory{})
	amlParser.AddInstructionFactory(instructions.ActivityInstructionFactory{})

	amlParser.AddForkJoinFactory(instructions.IfInstructionFactory{})
	amlParser.AddForkJoinFactory(instructions.ParallelInstructionFactory{})

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
