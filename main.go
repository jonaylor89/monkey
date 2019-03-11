package main

import (
	"flag"
	"fmt"
	"github.com/jonaylor89/monkey/compiler"
	"github.com/jonaylor89/monkey/evaluator"
	"github.com/jonaylor89/monkey/lexer"
	"github.com/jonaylor89/monkey/object"
	"github.com/jonaylor89/monkey/parser"
	"github.com/jonaylor89/monkey/repl"
	"github.com/jonaylor89/monkey/vm"
	"io/ioutil"
	"os"
	"os/user"
)

var engine = flag.String("engine", "vm", "use 'vm' or 'eval'")

func main() {

	if len(os.Args) < 2 {

		owner, err := user.Current()
		if err != nil {
			panic(err)
		}

		fmt.Printf("Hello %s! Welcome to interactive mode!\n",
			owner.Username)

		fmt.Printf("Enter monkey code\n")

		repl.Start(os.Stdin, os.Stdout)

	} else {
		dat, err := ioutil.ReadFile(os.Args[1])

		if err != nil {
			panic(err)
		}

		// env := object.NewEnvironment()
		// macroEnv := object.NewEnvironment()
		l := lexer.New(string(dat))
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			for _, msg := range p.Errors() {
				fmt.Println("\t" + msg + "\n")
			}

			os.Exit(1)
		}

		var result object.Object

		if *engine == "vm" {
			comp := compiler.New()
			err = comp.Compile(program)
			if err != nil {
				fmt.Printf("compiler error: %s", err)
			}

			machine := vm.New(comp.Bytecode())

			err = machine.Run()
			if err != nil {
				fmt.Printf("vm error: %s", err)
				return
			}

			result = machine.LastPoppedStackElem()
		} else {
			env := object.NewEnvironment()
			result = evaluator.Eval(program, env)
		}

		fmt.Println(result.Inspect())
	}
}
