package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/pverghese/gmon/repl"
)

func main() {

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s. This is the Monkey programming language!\n", user.Name)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
