package aml

import (
	"github.com/windler/awesm/aml/instructions"
	"github.com/windler/godepg/dotgraph"
)

type AMLFile struct {
	GraphType    string
	Instructions []instructions.AMLInstruction
}

func (af AMLFile) CreateDotGraph() *dotgraph.DotGraph {
	graph := dotgraph.New(af.GraphType)
	graph.SetEdgeGraphOptions(dotgraph.DotGraphOptions{
		"arrowhead": "open",
		"color":     "white",
		"fontcolor": "white",
		"splines":   "curved",
	})

	graph.SetNodeGraphOptions(dotgraph.DotGraphOptions{
		"fillcolor": "#336699",
		"style":     "filled",
		"fontcolor": "white",
		"fontname":  "Courier",
		"shape":     "oval",
	})

	graph.SetGraphOptions(dotgraph.DotGraphOptions{
		"bgcolor": "#333333",
	})

	for _, instruction := range af.Instructions {
		name := instruction.GetNodeName()
		predecessors := instruction.GetPredecessors()

		graph.AddNode(name)
		graph.AddNodeGraphPatternOptions(name, instruction.GetNodeOptions())

		for _, predecessor := range predecessors {
			edgeOptions := instruction.GetEdgeOptions()

			graph.AddDirectedEdge(predecessor.GetNodeName(), name, "")
			graph.AddEdgeGraphPatternOptions(name, edgeOptions)
		}
	}

	return graph
}
