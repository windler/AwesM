package main

import (
	"fmt"
	"os"

	"github.com/windler/dotgraph/renderer"

	"github.com/windler/awesm/aml"
	"github.com/windler/awesm/aml/constants"
	"github.com/windler/awesm/aml/instructions"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		fmt.Println("no file provided.")
		return
	}

	orientation := constants.TopDowm
	if len(argsWithoutProg) > 1 {
		orientation = constants.LeftRight
	}

	amlParser := aml.NewFileParser(argsWithoutProg[0])
	amlParser.AddInstructionFactory(instructions.StartInstructionFactory{})
	amlParser.AddInstructionFactory(instructions.EndInstructionFactory{})
	amlParser.AddInstructionFactory(instructions.ActivityInstructionFactory{})

	amlParser.AddForkJoinFactory(instructions.IfInstructionFactory{})
	amlParser.AddForkJoinFactory(instructions.ParallelInstructionFactory{
		Orientation: orientation,
	})

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

	graph := aml.CreateDotGraph(orientation)
	r.Render(graph.String())
}
