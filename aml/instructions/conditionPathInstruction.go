package instructions

import (
	"regexp"
	"strings"
)

type IfPathInstructionFactory struct{}

func (f IfPathInstructionFactory) New(name string, predecessor, parent *AMLInstruction) AMLInstruction {
	return AMLInstruction{
		Name:        strings.Replace(strings.Replace(name, "[", "", -1), "[", "", -1),
		Predecesors: []*AMLInstruction{parent},
		NodeOptions: map[string]string{
			"label": getNodeName(name),
		},
		EdgeOptions: map[string]string{
			"label": getLabelName(name),
		},
		PathBeginning: true,
		PathJoinNode:  parent.GetPathJoinNode(),
	}
}

func getLabelName(name string) string {
	r := regexp.MustCompile("\\[(.+)\\].+")
	return " [" + r.FindStringSubmatch(name)[1] + "]"
}

func getNodeName(name string) string {
	r := regexp.MustCompile("\\[.+\\](.+)")
	return r.FindStringSubmatch(name)[1]
}

func (f IfPathInstructionFactory) GetPattern() string {
	return "^\\[.+\\].+"
}
