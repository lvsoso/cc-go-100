package lexer

import (
	"fmt"
)

type TokenKind = int

const (
	TokenKindAdd = iota
	TokenKindSub
	TokenKindMul
	TokenKindDiv
	TokenKindNum
	TokenKindEof
)

type Token struct {
	Kind    TokenKind
	Value   int64
	Content string
}

type Lexer struct {
	sourceCode   string
	curChar      byte
	cursor       int
	CurrentToken *Token
}

func NewLexer(sourceCode string) *Lexer {
	lexer := &Lexer{
		sourceCode: sourceCode,
		// space ASCII 32
		curChar:      32,
		cursor:       0,
		CurrentToken: &Token{},
	}
	return lexer
}

func (lexer *Lexer) GetNextChar() {
	if lexer.cursor == len(lexer.sourceCode) {
		lexer.curChar = 0
		lexer.cursor++
	} else {
		lexer.curChar = lexer.sourceCode[lexer.cursor]
		lexer.cursor++
	}
}

func (lexer *Lexer) GetNextToken() {

	// space ASCII 32
	for lexer.curChar == 32 {
		lexer.GetNextChar()
	}

	var kind TokenKind
	value := 0
	startPos := lexer.cursor - 1
	// EOF
	if lexer.curChar == 0 {
		kind = TokenKindEof
		// add 43
	} else if lexer.curChar == '+' {
		kind = TokenKindAdd
		lexer.GetNextChar()
		// sub 45
	} else if lexer.curChar == '-' {
		kind = TokenKindSub
		lexer.GetNextChar()
		// mul 42
	} else if lexer.curChar == '*' {
		kind = TokenKindMul
		lexer.GetNextChar()
		// sub 47
	} else if lexer.curChar == '/' {
		kind = TokenKindDiv
		lexer.GetNextChar()
	} else if isDigit(lexer.curChar) {
		kind = TokenKindNum
		value = 0
		for {
			// 0 ASCII 48
			value = value*10 + int(lexer.curChar-48)
			lexer.GetNextChar()
			if !isDigit(lexer.curChar) {
				break
			}
		}
	} else {
		fmt.Printf("not support %c \n", lexer.curChar)
	}
	lexer.CurrentToken.Kind = kind
	lexer.CurrentToken.Value = int64(value)
	lexer.CurrentToken.Content = string(lexer.sourceCode[startPos : lexer.cursor-1])
}

func isDigit(b byte) bool {
	// 0 ASCII 48
	return (b-48 <= 9)
}
