package main

import (

)

func (l *Lexer) readNumber(firstRune rune) (string, error) {
	numberLiteral := ""
	decimal := false
	if firstRune != 0 { // firstRune is an optional param, this handles the case when firstRune is not declared in the call
		numberLiteral = append(numberLiteral, firstRune)
	}
	for {
		r, _, err := l.reader.ReadRune()
	}
}