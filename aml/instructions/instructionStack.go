package instructions

import (
	"github.com/golang-collections/collections/stack"
)

type InstructionStack struct {
	stack *stack.Stack
}

func NewInstructionStack() *InstructionStack {
	return &InstructionStack{
		stack: stack.New(),
	}
}

func (i InstructionStack) Push(item *AMLInstruction) {
	i.stack.Push(item)
}

func (i InstructionStack) Pop() *AMLInstruction {
	return i.stack.Pop().(*AMLInstruction)
}

func (i InstructionStack) Peek() *AMLInstruction {
	return i.stack.Peek().(*AMLInstruction)
}

func (i InstructionStack) Len() int {
	return i.stack.Len()
}
