package main

import (
	"fmt"
	"main/Monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n",
		user.Username)
	fmt.Printf("Feel free to tyoe in command\n")
	repl.Start(os.Stdin, os.Stdout)
}
