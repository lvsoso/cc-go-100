package main

import (
	"cc-go/lexer"
	"fmt"
)

var sourceCode = "  5 + 1-  3*4/2"

func main() {
	lex := lexer.NewLexer(sourceCode)
	for {
		lex.GetNextToken()
		if lex.CurrentToken.Kind == lexer.TokenKindEof {
			break
		}
		fmt.Println(lex.CurrentToken.Content)
	}
}
