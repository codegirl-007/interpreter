package repl

import (
	"bufio"
	"fmt"
	"io"
	"slang/lexer"
	"slang/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	// read source from input
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		// pass input to lexer
		line := scanner.Text()
		l := lexer.New(line)

		// print all the tokens the lexer gives until EOF
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
