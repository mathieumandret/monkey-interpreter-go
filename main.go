package main

import (
	"fmt"
	"interpreter/lexer/repl"
	"os"
)

func main() {
	fmt.Printf("REPL")
	repl.Start(os.Stdin, os.Stdout)
}
