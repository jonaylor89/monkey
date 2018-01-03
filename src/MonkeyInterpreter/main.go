
package main

import (
    "fmt"
    "os"
    "os/user"
    "MonkeyInterpreter/repl"
)

func main() { 
    user, err := user.Current()
    if err != nil { 
        panic(err) 
    }

    fmt.Printf("Hello %s! Welcome to interactive mode!\n",
                user.Username)

    fmt.Printf("Enter commands\n")

    repl.Start(os.Stdin, os.Stdout)
}
