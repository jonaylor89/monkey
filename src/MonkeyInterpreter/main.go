package main

import (
	"MonkeyInterpreter/repl"
    "MonkeyInterpreter/lexer"
    "MonkeyInterpreter/parser"
    "MonkeyInterpreter/object"
    "MonkeyInterpreter/evaluator"
	"fmt"
	"os"
	"os/user"
    "io/ioutil"
)

func main() {

    if len(os.Args) < 2 {
 
    	user, err := user.Current()
    	if err != nil {
    		panic(err)
    	}

    	fmt.Printf("Hello %s! Welcome to interactive mode!\n",
    		user.Username)

    	fmt.Printf("Enter commands\n")

    	repl.Start(os.Stdin, os.Stdout)
   
    }else {
        dat, err := ioutil.ReadFile(os.Args[1]) 

        if err != nil {
            panic(err) 
        }
        
        env := object.NewEnvironment()
        l := lexer.New(string(dat))
        p := parser.New(l)
        program := p.ParseProgram()

        if len(p.Errors()) != 0 {
            for _, msg := range p.Errors() {
                fmt.Println("\t" + msg + "\n") 
            } 

            os.Exit(1)
        }

        evaluated := evaluator.Eval(program, env)

        if evaluated != nil {
            fmt.Println(evaluated.Inspect()) 
        }

    }
}
