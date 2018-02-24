package instructions

import (
	"github.com/golang-collections/collections/stack"
)

//InstructionStack is a stack of AMLInstructions
type InstructionStack struct {
	stack *stack.Stack
}

//NewInstructionStack creates a new instructions stack
func NewInstructionStack() *InstructionStack {
	return &InstructionStack{
		stack: stack.New(),
	}
}

//Push pushes an item to the stack
func (i InstructionStack) Push(item *AMLInstruction) {
	i.stack.Push(item)
}

//Pop pops an item from the stack
func (i InstructionStack) Pop() *AMLInstruction {
	return i.stack.Pop().(*AMLInstruction)
}

//Peek retrieves the top most item
func (i InstructionStack) Peek() *AMLInstruction {
	return i.stack.Peek().(*AMLInstruction)
}

//Len gets the number of items on the stack
func (i InstructionStack) Len() int {
	return i.stack.Len()
}
