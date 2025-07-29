package main // temp

import (
	"bufio"
	"io"
)

type Token int

const (
	EOF Token = iota 	 // End of file (EOF) token, start value 0
	IDENTIFIER 	         // variable ids, value 1 etc.
	NUMBER			     // floats and ints
	ADD				     // +
	SUB				     // -
	MUL				     // *
	DIV				     // /
	MOD				     // %
	ASSIGN			     // =
	DOT				     // . (Currently just for use in floats)
	SEMICOLON		     // ;
)

var tokenLiterals = []string{
	EOF: "EOF",
	IDENTIFIER: "IDENTIFIER",
	NUMBER: "NUMBER",
	ADD: "ADD",
	SUB: "SUB",
	MUL: "MUL",
	DIV: "DIV",
	MOD: "MOD",
	ASSIGN: "ASSIGN",
	DOT: "DOT",
	SEMICOLON: "SEMICOLON",
}

// Print token method
func (t Token) String() string {
	return tokenLiterals[t]
}

type Position struct {
	row int
	col int
}

type Lexer struct {
	pos Position
	reader bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos: Position{row: 1, col: 0},
		reader: *bufio.NewReader(reader),
	}
}

func (l *Lexer) Lex() (Position, Token, string) {
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, EOF, ""
			}
			panic(err)
		}
		// state machine
		switch r {
		case '\n': // Special case where we need to reset to next line
			//x
		case '+':
			//y
		default:
			// logic for identifiers and numbers

		}

	}
}

func main() {

}