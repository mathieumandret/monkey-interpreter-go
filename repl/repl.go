package repl

import (
	"bufio"
	"fmt"
	"interpreter/lexer"
	"interpreter/token"
	"io"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.New(line)

		for tk := lexer.NextToken(); tk.Type != token.EOF; tk = lexer.NextToken() {
			fmt.Printf("%+v\n", tk)
		}

	}

}
