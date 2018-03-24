package aml

import (
	"errors"
	"io/ioutil"
	"os"
	"regexp"

	"gopkg.in/yaml.v2"

	"github.com/windler/awesm/aml/instructions"
)

//FileParser parses a yaml file
type FileParser struct {
	env          *amlEnv
	predecessor  *instructions.AMLInstruction
	parents      *instructions.InstructionStack
	currentLabel string
}

type amlEnv struct {
	file          string
	factories     []InstructionFactory
	forkFactories []ForkJoinFactory
}

//NewFileParser creates a new FileParser
func NewFileParser(file string) *FileParser {
	return &FileParser{
		env: &amlEnv{
			file:          file,
			factories:     []InstructionFactory{},
			forkFactories: []ForkJoinFactory{},
		},
		parents: instructions.NewInstructionStack(),
	}
}

//AddInstructionFactory adds a InstructionFactory
func (p *FileParser) AddInstructionFactory(factory InstructionFactory) {
	p.env.factories = append(p.env.factories, factory)
}

//AddForkJoinFactory adds a ForkJoinFactory
func (p *FileParser) AddForkJoinFactory(factory ForkJoinFactory) {
	p.env.forkFactories = append(p.env.forkFactories, factory)
}

//Parse parses a given AMLFile
func (p *FileParser) Parse() (File, error) {
	aml := File{}
	if _, err := os.Stat(p.env.file); err != nil {
		return aml, err
	}

	data, err := ioutil.ReadFile(p.env.file)
	if err != nil {
		return aml, err
	}

	var body map[string]interface{}
	yaml.Unmarshal(data, &body)

	p.parseDiagram(&aml, body["activity"])

	return aml, nil
}

func (p *FileParser) parseDiagram(aml *File, diagram interface{}) error {
	switch diagram.(type) {
	case []interface{}:
		for _, val := range diagram.([]interface{}) {
			if err := p.parseInstruction(val, aml); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *FileParser) parseInstruction(ins interface{}, aml *File) error {
	switch ins.(type) {
	case string:
		p.handleNewInstruction(aml, ins.(string))
	case map[interface{}]interface{}:
		for key, subDiagram := range ins.(map[interface{}]interface{}) {
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
	return nil
}

func (p *FileParser) handleNewInstruction(aml *File, ins string) {
	factory := p.getFactory(ins)

	if factory == nil {
		return
	}

	new := (*factory).New(ins)
	p.addInstruction(aml, new)
}

func (p *FileParser) handleNewFork(aml *File, fork string, subDiagram map[interface{}]interface{}) {
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

func (p *FileParser) addInstruction(aml *File, ins *instructions.AMLInstruction) {
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

func (p *FileParser) getCurrentParent() *instructions.AMLInstruction {
	if p.parents.Len() > 0 {
		parent := p.parents.Peek()
		return parent
	}

	return nil
}

func (p *FileParser) getFactory(line string) *InstructionFactory {
	for _, f := range p.env.factories {
		if matches, _ := regexp.MatchString(f.GetPattern(), line); matches {
			return &f
		}
	}

	return nil
}

func (p *FileParser) getForkJoinFactory(line string) *ForkJoinFactory {
	for _, f := range p.env.forkFactories {
		if matches, _ := regexp.MatchString(f.GetPattern(), line); matches {
			return &f
		}
	}

	return nil
}
