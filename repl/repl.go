package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/DougAlves/interpreter/lexer"
	"github.com/DougAlves/interpreter/parser"
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
		p := parser.New(l)
		program := p.ParseProgram()
		if errors := p.Errors(); len(errors) > 0 {
			printParserErros(out, errors)
		}
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErros(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
