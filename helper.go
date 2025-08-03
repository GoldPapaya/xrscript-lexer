package main

import (
	"io"
	"unicode"
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