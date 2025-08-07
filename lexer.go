package main // temp

import (
	"fmt"
	"io"
	"unicode"
)

func (l *Lexer) Lex() Token {
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
				return Token{pos: Position{l.pos.row, l.pos.col - len(numberLiteral)}, typ: 2, literal: numberLiteral}
			} else { // *Currently just a catch-all for non-numbers. Will need to change later if adding strings & more
				if unicode.IsLetter(r) {

				}
			}
		}

	}

}

func main() {
	fmt.Println(EOF)
}
