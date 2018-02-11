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
	parentJoinNode  map[string]*instructions.AMLInstruction
}

func NewFileParser(file string) *AMLFileParser {
	return &AMLFileParser{
		file:            file,
		factories:       []AMLInstructionFactory{},
		parents:         instructions.NewInstructionStack(),
		parentPathNodes: map[string][]instructions.InstructionStack{},
		parentJoinNode:  map[string]*instructions.AMLInstruction{},
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

	strippedLine := strings.Replace(line, ".", "", -1)
	factory := p.getFactory(strippedLine)

	new := (*factory).New(strippedLine)

	p.handleFork(new, factory, line, aml)
	p.handleParents(line, aml)
	p.addInstruction(aml, new)
}

func (p *AMLFileParser) addInstruction(aml *AMLFile, ins *instructions.AMLInstruction) {
	if p.getCurrentParent() != nil {
		pathStacks := p.parentPathNodes[p.getCurrentParent().Name]
		currentStack := pathStacks[len(pathStacks)-1]
		currentStack.Push(ins)
	}

	if p.predecessor != nil {
		ins.Predecessors = append(ins.Predecessors, p.predecessor)
	}
	aml.Instructions = append(aml.Instructions, *ins)
	p.predecessor = ins
}

func (p *AMLFileParser) handlePathBeginningNode(ins *instructions.AMLInstruction) {
	parent := p.getCurrentParent()

	p.parentPathNodes[parent.Name] = append(p.parentPathNodes[parent.Name], *instructions.NewInstructionStack())
	p.predecessor = parent
}

func (p *AMLFileParser) handleFork(new *instructions.AMLInstruction, factory *AMLInstructionFactory, line string, aml *AMLFile) {
	forkNode := (*factory).NewForkNode(line)
	if forkNode != nil {
		tabs := strings.Count(line, ".") / 2

		if p.parents.Len() < tabs {
			p.addInstruction(aml, forkNode)
			p.parents.Push(forkNode)
			p.parentPathNodes[forkNode.Name] = []instructions.InstructionStack{
				*instructions.NewInstructionStack(),
			}

			joinNode := (*factory).NewJoinNode(line, forkNode)
			p.parentJoinNode[forkNode.Name] = joinNode
		}

		p.handlePathBeginningNode(new)
	}
}

func (p *AMLFileParser) handleParents(line string, aml *AMLFile) {
	tabs := strings.Count(line, ".") / 2

	for tabs < p.parents.Len() {
		lastParent := p.parents.Pop()
		parentJoinNode := p.parentJoinNode[lastParent.Name]

		for _, stack := range p.parentPathNodes[lastParent.Name] {
			if stack.Len() > 0 {
				lastPathElem := stack.Peek()
				if lastPathElem != nil {
					parentJoinNode.Predecessors = append(parentJoinNode.Predecessors, lastPathElem)
				}
			}
		}

		p.addInstruction(aml, parentJoinNode)
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
