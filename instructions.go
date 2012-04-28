package interpreter

type Instruction interface {
	GetType() InstructionType
	String() string
}

// a non-control instruction. Should be used to perform
// data manipulation in your domain.
type DataInstruction interface {
	Instruction
	Execute()
}

// causes current execution point to be pushed onto the stack
// and execution moves to the first instruction in the named
// by GetBlockName(). Execution will return to the current point
// when the called block completes.
type CallInstruction interface {
	Instruction
	GetBlockName() string
}

// moves execution to the requested step in the current block
type JumpRelativeInstruction interface {
	Instruction
	// zero based
	GetNextStepNumber() int
}

// splits execution. One branch continues from the current
// execution point. Simultaneously, a new interpreter is
// started with start point being the current block at
// the requested step
type SplitRelativeInstruction interface {
	JumpRelativeInstruction
}
