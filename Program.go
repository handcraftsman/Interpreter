package interpreter

// your program must implement this interface
type Program interface {
	GetBlock(name string) []Instruction
}
