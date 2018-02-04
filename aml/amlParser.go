package aml

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/windler/awesm/aml/instructions"
)

type AMLFileParser struct {
	file            string
	factories       []AMLInstructionFactory
	predecessor     *instructions.AMLInstruction
	parents         *instructions.InstructionStack
	parentPathNodes map[string][]instructions.InstructionStack
}

func NewFileParser(file string) *AMLFileParser {
	return &AMLFileParser{
		file:            file,
		factories:       []AMLInstructionFactory{},
		parents:         instructions.NewInstructionStack(),
		parentPathNodes: map[string][]instructions.InstructionStack{},
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

	for _, line := range strings.Split(string(data[:]), "\n") {
		p.parseLine(&aml, line)
	}

	return aml, nil
}

func (p *AMLFileParser) parseLine(aml *AMLFile, line string) {
	if strings.TrimSpace(line) == "" {
		return
	}

	p.handleParents(line, aml)

	parent := p.getCurrentParent()
	strippedLine := strings.Replace(line, ".", "", -1)
	factory := p.getFactory(strippedLine)

	new := (*factory).New(strippedLine, p.predecessor, parent)
	p.handlePathBeginningNode(new)

	p.addInstruction(aml, &new)
}

func (p *AMLFileParser) addInstruction(aml *AMLFile, ins *instructions.AMLInstruction) {
	if p.getCurrentParent() != nil {
		pathStacks := p.parentPathNodes[p.getCurrentParent().Name]
		currentStack := pathStacks[len(pathStacks)-1]
		currentStack.Push(ins)
	}

	aml.Instructions = append(aml.Instructions, *ins)
	p.predecessor = ins
}

func (p *AMLFileParser) handlePathBeginningNode(ins instructions.AMLInstruction) {
	if ins.IsPathBeginning() {
		parent := p.getCurrentParent()

		p.parentPathNodes[parent.Name] = append(p.parentPathNodes[parent.Name], *instructions.NewInstructionStack())
	}
}

func (p *AMLFileParser) handleParents(line string, aml *AMLFile) {
	tabs := strings.Count(line, ".") / 2

	//new parent
	if p.parents.Len() < tabs {
		p.parents.Push(p.predecessor)
		p.parentPathNodes[(*p.predecessor).Name] = []instructions.InstructionStack{
			*instructions.NewInstructionStack(),
		}
	}

	//parent left
	for tabs < p.parents.Len() {
		lastParent := p.parents.Pop()

		for _, stack := range p.parentPathNodes[lastParent.Name] {
			if stack.Len() > 0 {
				lastPathElem := stack.Peek()
				lastParent.GetPathJoinNode().Predecesors = append(lastParent.GetPathJoinNode().Predecesors, lastPathElem)
			}
		}

		p.addInstruction(aml, lastParent.GetPathJoinNode())
	}
}

func (p *AMLFileParser) getCurrentParent() *instructions.AMLInstruction {
	if p.parents.Len() > 0 {
		parent := p.parents.Peek()
		return parent
	}

	return nil
}

func (p *AMLFileParser) getFactory(line string) *AMLInstructionFactory {
	for _, f := range p.factories {
		if matches, _ := regexp.MatchString(f.GetPattern(), line); matches {
			return &f
		}
	}

	return nil
}
