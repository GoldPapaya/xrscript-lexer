package main

import (
	"io"
	"unicode"
	"bufio"
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

func (l *Lexer) readNumber(firstRune rune) (string, error) {
	var numRunes []rune
	decimal := false
	if firstRune != 0 { // firstRune is an optional param, this handles the case when firstRune is not declared in the call
		numRunes = append(numRunes, firstRune)
	}
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return string(numRunes), nil
			}
			panic(err)
		}
		l.pos.col += 1
		if !decimal && r == '.' {
			decimal = true
			numRunes = append(numRunes, r)
			continue
		}
		if !unicode.IsDigit(r) {
			err = l.reader.UnreadRune()
			if err != nil {
				return "", err
			}
			l.pos.col -= 1
			return string(numRunes), nil
		}
		numRunes = append(numRunes, r)
	}
}