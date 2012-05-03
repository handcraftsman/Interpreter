package interpreter

import "fmt"

type Interpreter struct {
	program             Program
	stack               []blockIndex
	stepCount           int
	maxSteps            int
	missingBlockHandler func(blockName string) *[]Instruction
	haltExecution       func() bool
}

func NewInterpreter(p Program) *Interpreter {
	in := Interpreter{
		program: p,
		stack:   make([]blockIndex, 0, 10),
	}
	return in.
		WithMaxSteps(0).
		WithMissingBlockHandler(func(blockName string) *[]Instruction {
		panic("failed to find block named '" + blockName + "'")
	})
}

// limit the number of instructions to execute, defaults to Int32Max
func (in *Interpreter) WithMaxSteps(maxSteps int) *Interpreter {
	if maxSteps < 0 {
		panic("maxSteps must be >= 0. 0 means 'until completion'")
	}
	in.maxSteps = maxSteps
	if maxSteps == 0 {
		in.haltExecution = func() bool { return false }
	} else {
		in.haltExecution = func() bool { return in.stepCount >= in.maxSteps }
	}

	return in
}

// provide an arbitrary halt condition. May be called multiple times.
func (in *Interpreter) WithHaltIf(haltIf func() bool) *Interpreter {
	f := in.haltExecution
	in.haltExecution = func() bool { return haltIf() || f() }

	return in
}

// by default the system will panic if a call instruction requests a non-existent block
// override that behavior with this function
func (in *Interpreter) WithMissingBlockHandler(missingBlockHandler func(blockName string) *[]Instruction) *Interpreter {
	in.missingBlockHandler = missingBlockHandler
	return in
}

func (in *Interpreter) Run(startingBlockName string, args CallArgs, startAt int) {
	instructions := in.getNamedBlock(startingBlockName, args)
	in.pushState(startingBlockName, instructions, startAt)

	for {
		block := in.popState()
		for {
			if block.instructions == nil {
				return
			}
			if block.index < len(*block.instructions) {
				break
			}
			block = in.popState()
		}

		blockName := block.name
		instructions := *block.instructions
		index := block.index

		for i := index; i < len(instructions) && !in.haltExecution(); i, in.stepCount = i+1, in.stepCount+1 {
			instr := instructions[i]
			switch instr.GetType() {
			case Data:
				dataInstr := instr.(DataInstruction)
				dataInstr.Execute()
			case Call:
				callInstr := instr.(CallInstruction)
				in.pushState(blockName, &instructions, i+1)
				blockName = callInstr.GetBlockName()
				args = callInstr.GetArgs()
				instructions = *in.getNamedBlock(blockName, args)
				i = -1
			case Jump:
				jumpInstr := instr.(JumpRelativeInstruction)
				i = i - 1 + jumpInstr.GetNextStepNumber()
			case Split:
				splitInstr := instr.(SplitRelativeInstruction)
				other := NewInterpreter(in.program)
				other.missingBlockHandler = in.missingBlockHandler
				other.haltExecution = in.haltExecution
				go other.Run(blockName, args, i-1+splitInstr.GetNextStepNumber())
			default:
				panic(fmt.Sprint("don't know how to handle instruction type '", instr.GetType(), "'"))
			}
		}
	}
}

func (in *Interpreter) pushState(name string, instructions *[]Instruction, index int) {
	in.stack = append(in.stack, blockIndex{name: name, instructions: instructions, index: index})
}

func (in *Interpreter) popState() blockIndex {
	if len(in.stack) == 0 {
		return *new(blockIndex)
	}
	block := in.stack[len(in.stack)-1]
	in.stack = in.stack[:len(in.stack)-1]

	return block
}

func (in *Interpreter) getNamedBlock(blockName string, args CallArgs) *[]Instruction {
	instructions := in.program.GetBlock(blockName, args)
	if instructions == nil {
		return in.missingBlockHandler(blockName)
	}
	return &instructions
}

type blockIndex struct {
	name         string
	instructions *[]Instruction
	index        int
}
