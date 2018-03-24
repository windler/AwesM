package aml

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/windler/awesm/aml/instructions"
	"github.com/windler/awesm/aml/mocks"
)

func TestNoFactories(t *testing.T) {
	file := createTestFile(`
activity:
  - start`)

	parser := NewFileParser(file)

	amlFile, err := parser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 0, len(amlFile.Instructions))
}

func TestFactoryPatternDoesNotMatch(t *testing.T) {
	file := createTestFile(`
activity:
  - start`)

	factoryMock := &mocks.InstructionFactory{}
	factoryMock.On("GetPattern").Return("special")

	parser := NewFileParser(file)
	parser.AddInstructionFactory(factoryMock)

	amlFile, err := parser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 0, len(amlFile.Instructions))
	factoryMock.AssertExpectations(t)
}

func TestForkNode(t *testing.T) {
	file := createTestFile(`
activity:
  - start
  - fork:
      forkPath1:
        - sub
  - end`)

	testInstruction := &instructions.AMLInstruction{
		Name:        "testName",
		NodeOptions: make(map[string]string),
		EdgeOptions: make(map[string]string),
	}

	forkInstruction := &instructions.AMLInstruction{
		Name:        "fork",
		NodeOptions: make(map[string]string),
		EdgeOptions: make(map[string]string),
	}

	joinInstruction := &instructions.AMLInstruction{
		Name:        "join",
		NodeOptions: make(map[string]string),
		EdgeOptions: make(map[string]string),
	}

	factoryMock := &mocks.InstructionFactory{}
	factoryMock.On("GetPattern").Return(".*")
	factoryMock.On("New", "start").Return(testInstruction)
	factoryMock.On("New", "end").Return(testInstruction)
	factoryMock.On("New", "sub").Return(testInstruction)

	forkFactoryMock := &mocks.ForkJoinFactory{}
	forkFactoryMock.On("GetPattern").Return("fork")
	forkFactoryMock.On("ProvidesPathLabels").Return(true)
	forkFactoryMock.On("NewForkNode", "fork").Return(forkInstruction)
	forkFactoryMock.On("NewJoinNode", "fork", forkInstruction).Return(joinInstruction)

	parser := NewFileParser(file)
	parser.AddForkJoinFactory(forkFactoryMock)
	parser.AddInstructionFactory(factoryMock)

	amlFile, err := parser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 5, len(amlFile.Instructions))
	factoryMock.AssertExpectations(t)
}

func createTestFile(content string) string {
	file, _ := ioutil.TempFile("", "aml_test")
	ioutil.WriteFile(file.Name(), []byte(content), 0644)

	return file.Name()
}
