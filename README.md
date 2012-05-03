# Interpreter - simple domain independent interpreter

This is an interpreter that can be used to handle the flow control of a program that implements its interfaces.

## Usage

Interpreter is compatible with Go 1. Add it to your package repository:

	go get "github.com/handcraftsman/Interpreter"

then use it in your program:

	import interpreter "github.com/handcraftsman/Interpreter"

create your program. It must implement the Program interface and your instructions must implement one of the supported instruction interfaces.
	
	DataInstruction - perform data manipulation in your domain
	CallInstruction - call a named subroutine
	JumpRelativeInstruction - move execution forward or backward from the current execution point
	SplitRelativeInstruction - splits off a new Interpreter that starts executing in the current block at an offset you provide
	
create an interpreter for your program

	in := interpreter.NewInterpreter(program)
	
configuration methods for the interpreter include:
	
	WithMaxSteps(int) - limit the number of instructions to execute, defaults to Int32Max
	WithHaltIf(func() bool) - an arbitrary halt condition
	WithMissingBlockHandler(func(blockName string)*[]Instruction) - by default the system will panic if a call instruction requests a non-existent block

e.g.
	
	in := interpreter.NewInterpreter(program).
		WithMaxSteps(1000)
	
run the interpreteter by passing in the name of the starting block, its args, and the index of first instruction to run in that block

	in.Run("main", nil, 0)

	
## License		

[MIT License](http://www.opensource.org/licenses/mit-license.php)
