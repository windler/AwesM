package aml

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/windler/awesm/aml/instructions"
	"github.com/windler/awesm/aml/mocks"
)

func TestNoFactories(t *testing.T) {
	file := createTestFile(`#start`)

	parser := NewFileParser(file)

	amlFile, err := parser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 0, len(amlFile.Instructions))
}

func TestFactoryPatternDoesNotMatch(t *testing.T) {
	file := createTestFile(`#start`)

	factoryMock := &mocks.AMLInstructionFactory{}
	factoryMock.On("GetPattern").Return("^!.*")

	parser := NewFileParser(file)
	parser.AddFactory(factoryMock)

	amlFile, err := parser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 0, len(amlFile.Instructions))
	factoryMock.AssertExpectations(t)
}

func TestFactoryPatternMatchesNoForkNode(t *testing.T) {
	file := createTestFile(`#start`)

	testInstruction := &instructions.AMLInstruction{
		Name:        "testName",
		NodeOptions: make(map[string]string),
		EdgeOptions: make(map[string]string),
	}

	factoryMock := &mocks.AMLInstructionFactory{}
	factoryMock.On("GetPattern").Return("^#.*")
	factoryMock.On("New", "#start").Return(testInstruction)
	factoryMock.On("NewForkNode", "#start").Return(nil)

	parser := NewFileParser(file)
	parser.AddFactory(factoryMock)

	amlFile, err := parser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(amlFile.Instructions))
	assert.Equal(t, *testInstruction, amlFile.Instructions[0])
	factoryMock.AssertExpectations(t)
}

func TestForkNode(t *testing.T) {
	file := createTestFile(`
#start
..!fork
....#forkPath1
#end`)

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

	factoryMock := &mocks.AMLInstructionFactory{}
	factoryMock.On("GetPattern").Return("^#.*")
	factoryMock.On("New", "#start").Return(testInstruction)
	factoryMock.On("NewForkNode", "#start").Return(nil)
	factoryMock.On("New", "#forkPath1").Return(testInstruction)
	factoryMock.On("NewForkNode", "#forkPath1").Return(nil)
	factoryMock.On("New", "#end").Return(testInstruction)
	factoryMock.On("NewForkNode", "#end").Return(nil)

	forkFactoryMock := &mocks.AMLInstructionFactory{}
	forkFactoryMock.On("GetPattern").Return("!.*")
	forkFactoryMock.On("New", "!fork").Return(testInstruction)
	forkFactoryMock.On("NewForkNode", "!fork").Return(forkInstruction)
	forkFactoryMock.On("NewJoinNode", "!fork", forkInstruction).Return(joinInstruction)

	parser := NewFileParser(file)
	parser.AddFactory(forkFactoryMock)
	parser.AddFactory(factoryMock)

	amlFile, err := parser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 6, len(amlFile.Instructions))
	factoryMock.AssertExpectations(t)
}

func TestForkNodeNoHirachy(t *testing.T) {
	file := createTestFile(`
#start
!fork
..#forkPath1
#end`)

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

	factoryMock := &mocks.AMLInstructionFactory{}
	factoryMock.On("GetPattern").Return("^#.*")
	factoryMock.On("New", "#start").Return(testInstruction)
	factoryMock.On("NewForkNode", "#start").Return(nil)

	forkFactoryMock := &mocks.AMLInstructionFactory{}
	forkFactoryMock.On("GetPattern").Return("!.*")
	forkFactoryMock.On("New", "!fork").Return(testInstruction)
	forkFactoryMock.On("NewForkNode", "!fork").Return(forkInstruction)

	parser := NewFileParser(file)
	parser.AddFactory(forkFactoryMock)
	parser.AddFactory(factoryMock)

	amlFile, err := parser.Parse()

	assert.Nil(t, err)
	assert.Equal(t, 0, len(amlFile.Instructions))
	factoryMock.AssertExpectations(t)
}

func createTestFile(content string) string {
	file, _ := ioutil.TempFile("", "aml_test")
	ioutil.WriteFile(file.Name(), []byte(content), 0644)

	return file.Name()
}
