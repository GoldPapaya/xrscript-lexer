package main // temp

import (

)

type TokenType int

const (
	EOF TokenType = iota // End of file (EOF) token, start value 0
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

var tokens = []string{
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
func (t TokenType) String() string {
	return tokens[t]
}