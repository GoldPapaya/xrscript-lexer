package main

import (
	"io"
)

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
		// some fxn that checks whether the rune is a number. this will handle invalid input like two decimals in one number.
		// at the end reminder to do the base case append num to numRUnes
	}
}