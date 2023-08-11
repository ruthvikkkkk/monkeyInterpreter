package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nHello %s, Welcome to to the Monkey Programming Language!\n", user.Username)
	fmt.Println("Go Crazy!")
	repl.Start(os.Stdin, os.Stdout)
}
