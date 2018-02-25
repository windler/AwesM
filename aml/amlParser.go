package aml

import (
	"errors"
	"io/ioutil"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"

	"github.com/windler/awesm/aml/instructions"
)

type AMLFileParser struct {
	file          string
	factories     []InstructionFactory
	forkFactories []ForkJoinFactory
	predecessor   *instructions.AMLInstruction
	parents       *instructions.InstructionStack
	currentLabel  string
}

func NewFileParser(file string) *AMLFileParser {
	return &AMLFileParser{
		file:          file,
		factories:     []InstructionFactory{},
		forkFactories: []ForkJoinFactory{},
		parents:       instructions.NewInstructionStack(),
	}
}

func (p *AMLFileParser) AddInstructionFactory(factory InstructionFactory) {
	p.factories = append(p.factories, factory)
}

func (p *AMLFileParser) AddForkJoinFactory(factory ForkJoinFactory) {
	p.forkFactories = append(p.forkFactories, factory)
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

	var body map[string]interface{}
	yaml.Unmarshal(data, &body)

	p.parseDiagram(&aml, body["diagram"])

	return aml, nil
}

func (p *AMLFileParser) parseDiagram(aml *AMLFile, diagram interface{}) error {
	switch diagram.(type) {
	case []interface{}:
		for _, val := range diagram.([]interface{}) {
			switch val.(type) {
			case string:
				p.handleNewInstruction(aml, val.(string))
			case map[interface{}]interface{}:
				for key, subDiagram := range val.(map[interface{}]interface{}) {
					switch subDiagram.(type) {
					case map[interface{}]interface{}:
						p.handleNewFork(aml, key.(string), subDiagram.(map[interface{}]interface{}))
					default:
						return errors.New("Invalid YAML after " + key.(string))
					}
				}
			default:
				return errors.New("Invalid YAML given")
			}
		}
	}

	return nil
}

func (p *AMLFileParser) handleNewInstruction(aml *AMLFile, ins string) {
	factory := p.getFactory(ins)

	if factory == nil {
		return
	}

	new := (*factory).New(ins)
	p.addInstruction(aml, new)
}

func (p *AMLFileParser) handleNewFork(aml *AMLFile, fork string, subDiagram map[interface{}]interface{}) {
	factory := p.getForkJoinFactory(fork)

	forkNode := (*factory).NewForkNode(fork)
	joinNode := (*factory).NewJoinNode(fork, forkNode)

	lastPathNodes := []*instructions.AMLInstruction{}

	p.addInstruction(aml, forkNode)

	for label, diagram := range subDiagram {
		if (*factory).ProvidesPathLabels() {
			p.currentLabel = label.(string)
		}

		p.parents.Push(forkNode)
		p.parseDiagram(aml, diagram)
		lastPathNodes = append(lastPathNodes, &aml.Instructions[len(aml.Instructions)-1])
	}

	joinNode.Predecessors = append(joinNode.Predecessors, lastPathNodes...)
	p.addInstruction(aml, joinNode)
}

func (p *AMLFileParser) addInstruction(aml *AMLFile, ins *instructions.AMLInstruction) {
	if p.getCurrentParent() != nil {
		ins.Predecessors = append(ins.Predecessors, p.getCurrentParent())
	} else if p.predecessor != nil {
		ins.Predecessors = append(ins.Predecessors, p.predecessor)
	}
	aml.Instructions = append(aml.Instructions, *ins)
	p.predecessor = ins

	if p.currentLabel != "" {
		ins.EdgeOptions["label"] = p.currentLabel
		p.currentLabel = ""
	}

	if p.parents.Len() > 0 {
		p.parents.Pop()
	}
}

func (p *AMLFileParser) getCurrentParent() *instructions.AMLInstruction {
	if p.parents.Len() > 0 {
		parent := p.parents.Peek()
		return parent
	}

	return nil
}

func (p *AMLFileParser) getFactory(line string) *InstructionFactory {
	for _, f := range p.factories {
		if matches, _ := regexp.MatchString(f.GetPattern(), line); matches {
			return &f
		}
	}

	return nil
}

func (p *AMLFileParser) getForkJoinFactory(line string) *ForkJoinFactory {
	for _, f := range p.forkFactories {
		if matches, _ := regexp.MatchString(f.GetPattern(), line); matches {
			return &f
		}
	}

	return nil
}
