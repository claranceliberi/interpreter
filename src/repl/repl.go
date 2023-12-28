package repl

import (
	"bufio"
	"fmt"
	"go/token"
	"io"

	"github.com/claranceliberi/monkey-interpreter/waiig_code_1.3/01/src/monkey/lexer"
)



const PROMPT = ">> "

func Start(in io.Reader, out io.Writer){
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if(!scanned){
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = tok.NextToken(){
			fmt.Printf("%+v\n", tok)
		}
	}
}