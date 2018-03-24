package aml

import (
	"github.com/windler/awesm/aml/instructions"
	"github.com/windler/dotgraph/graph"
)

//File represents an *.aml file
type File struct {
	GraphType    string
	Instructions []instructions.AMLInstruction
}

//CreateDotGraph creates a graph.dograph based on the *.aml file
func (af File) CreateDotGraph(orientation string) *graph.DotGraph {
	g := graph.New(af.GraphType)
	g.SetEdgeGraphOptions(graph.DotGraphOptions{
		"arrowhead": "open",
		"color":     "white",
		"fontcolor": "white",
		"splines":   "curved",
	})

	g.SetNodeGraphOptions(graph.DotGraphOptions{
		"fillcolor": "#336699",
		"style":     "filled",
		"fontcolor": "white",
		"fontname":  "Courier",
		"shape":     "oval",
	})

	g.SetGraphOptions(graph.DotGraphOptions{
		"bgcolor": "#333333",
		"rankdir": orientation,
	})

	for _, instruction := range af.Instructions {
		name := instruction.GetNodeName()
		predecessors := instruction.GetPredecessors()

		g.AddNode(name)
		g.AddNodeGraphPatternOptions(name, instruction.GetNodeOptions())

		for _, predecessor := range predecessors {
			edgeOptions := instruction.GetEdgeOptions()

			g.AddDirectedEdge(predecessor.GetNodeName(), name, "")
			g.AddEdgeGraphPatternOptions(name, edgeOptions)
		}
	}

	return g
}
