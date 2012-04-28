package interpreter

type InstructionType int

// supported instruction types
const (
	Data InstructionType = 1 + iota
	Call
	Jump
	Split
)
