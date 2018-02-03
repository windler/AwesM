package aml

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/golang-collections/collections/stack"
	"github.com/windler/awesm/aml/instructions"
)

type AMLFileParser struct {
	file      string
	factories []AMLInstructionFactory
}

func NewFileParser(file string) *AMLFileParser {
	return &AMLFileParser{
		file:      file,
		factories: []AMLInstructionFactory{},
	}
}

func (p *AMLFileParser) AddFactory(factory AMLInstructionFactory) {
	p.factories = append(p.factories, factory)
}

func (p *AMLFileParser) Parse() (AMLFile, error) {
	aml := AMLFile{}
	if _, err := os.Stat(p.file); err != nil {
		return aml, err
	}

	data, err := ioutil.ReadFile(p.file)
	if err != nil {
		return aml, err
	}

	var predecessor *instructions.AMLInstruction
	parents := stack.New()
	for _, line := range strings.Split(string(data[:]), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		strippedLine := strings.Replace(line, ".", "", -1)
		for _, factory := range p.factories {
			if matches, _ := regexp.MatchString(factory.GetPattern(), strippedLine); matches {
				tabs := strings.Count(line, ".") / 2
				fmt.Println(tabs)

				var parent instructions.AMLInstruction

				if parents.Len() < tabs {
					parents.Push(*predecessor)
				}
				if parents.Len() > 0 {
					parent = parents.Peek().(instructions.AMLInstruction)
				}
				for tabs < parents.Len() {
					parent = parents.Pop().(instructions.AMLInstruction)
					parent.GetPathJoinNode().Predecesors = append(parent.GetPathJoinNode().Predecesors, predecessor)
					aml.Instructions = append(aml.Instructions, *parent.GetPathJoinNode())
					predecessor = parent.GetPathJoinNode()
				}

				new := factory.New(strippedLine, predecessor, &parent)
				aml.Instructions = append(aml.Instructions, new)
				if new.IsPathBeginning() && predecessor.GetPathJoinNode() == nil {
					parent.GetPathJoinNode().Predecesors = append(parent.GetPathJoinNode().Predecesors, predecessor)
				}

				predecessor = &new
				break
			}
		}
	}

	return aml, nil
}
