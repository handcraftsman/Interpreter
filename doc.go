// a simple domain independent interpreter
//
// This is an interpreter that can be used to handle the flow control of a program that implements its interfaces.
//
// Usage
//
// Interpreter is compatible with Go 1. Add it to your package repository:
//
//	go get "github.com/handcraftsman/Interpreter"
//
// then use it in your program:
//
//	import interpreter "github.com/handcraftsman/Interpreter"
//
// create your program. It must implement the Program interface and your instructions must implement one of the supported instruction interfaces.
// 	
// 	DataInstruction
// 	CallInstruction
// 	JumpRelativeInstruction
// 	SplitRelativeInstruction
// 	
// create an interpreter for your program
// 
// 	in := interpreter.NewInterpreter(program)
// 	
// configuration methods for the interpreter include:
// 	
// 	WithMaxSteps(int)
// 	WithHaltIf(func() bool)
// 	WithMissingBlockHandler(func(blockName string)*[]Instruction)
// 
// e.g.
// 	
// 	in := interpreter.NewInterpreter(program).
// 		WithMaxSteps(1000)
// 	
// run the interpreteter by passing in the name of the starting block, its args, and the index of first instruction to run in that block
// 
// 	in.Run("main", nil, 0)
package interpreter
