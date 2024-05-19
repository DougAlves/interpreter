package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/DougAlves/interpreter/lexer"
	"github.com/DougAlves/interpreter/parser"
	"github.com/DougAlves/interpreter/token"
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
		l := lexer.New(line)
		parser.New(l)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

