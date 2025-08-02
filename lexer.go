package main // temp

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

type TokenType int

const (
	EOF TokenType = iota 	 // End of file (EOF) token, start value 0
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

// Print token method, mostly for testing
func (t TokenType) String() string {
	return tokenLiterals[t]
}

type Position struct {
	row int
	col int
}

type Token struct {
	pos Position
	typ TokenType
	literal string
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

func (l *Lexer) newline() {
	l.pos.row += 1
	l.pos.col = 0
}

func (l *Lexer) Lex() (Token) {
	for {
		r, _, err := l.reader.ReadRune() // .ReadRune() returns Rune, size, err - we only care about the first and last
		if err != nil {
			if err == io.EOF {
				return Token{pos: l.pos, typ: 0, literal: ""}
			}
			panic(err)
		}
		// State machine
		switch r {
		case '\n': // Special case where we need to reset to next line
			l.newline()
		case '+':
			return Token{pos: l.pos, typ: 3, literal: "+"}
		case '-': // TODO This one is way more annoying than it looks since the grammar handles negative numbers. implement lookahead or smth
			return Token{pos: l.pos, typ: 4, literal: "-"}
		case '*':
			return Token{pos: l.pos, typ: 5, literal: "*"}
		case '/':
			return Token{pos: l.pos, typ: 6, literal: "/"}
		case '%':
			return Token{pos: l.pos, typ: 7, literal: "%"}
		case '=':
			return Token{pos: l.pos, typ: 8, literal: "="}
		case '.':
			return Token{pos: l.pos, typ: 9, literal: "."}
		case ';':
			return Token{pos: l.pos, typ: 10, literal: ";"}
		default:
			// logic for identifiers and numbers
			if unicode.IsDigit(r) {
				numberLiteral, err := l.readNumber(r)
				if err != nil {
					panic(err)
				}
				return Token{pos: Position{l.pos.row, l.pos.col-len(numberLiteral)}, typ: 2, literal: numberLiteral}
			} else { // *Currently just a catch-all for non-numbers. Will need to change later if adding strings & more
				for unicode.IsLetter(r) {

				}
			}


		}

	}
}

func main() {
	fmt.Println(EOF)
}