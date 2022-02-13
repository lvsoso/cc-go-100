package lexer

import (
	"fmt"
	"testing"
)

var sourceCode = "  5 + 1-  3*4/2"

func Test_Lexer(t *testing.T) {
	lex := NewLexer(sourceCode)
	for {
		lex.GetNextToken()
		if lex.CurrentToken.Kind == TokenKindEof {
			break
		}
		fmt.Println(lex.CurrentToken.Content)
	}
}
