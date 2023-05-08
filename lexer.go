package main

import (
	"fmt"
	"unicode"
)

type Lexer struct {
	expression []string
	position   int
	tokens     []Token
}

func NewLexer(expr []string) *Lexer {
	return &Lexer{
		expression: expr,
		position:   0,
		tokens:     []Token{},
	}
}

func (l *Lexer) Scan() []Token {
	for l.position < len(l.expression) {
		switch {
		case l.isNumber():
			l.scanNumber()
			l.next()
		case l.isOperation():
			l.scanOperation()
			l.next()
		default:
			panic(fmt.Sprintf("unexpected character at position %d: %s", l.position, l.peek()))
		}
	}
	return l.tokens
}

func (l *Lexer) scanNumber() {
	word := l.expression[l.position]

	l.tokens = append(l.tokens, Token{_type: INT, Value: word})
}

func (l *Lexer) isNumber() bool {
	for _, l := range l.peek() {
		if !unicode.IsLetter(l) {
			return false
		}
	}

	if _, ok := digits[l.peek()]; ok {
		return true
	}

	return false
}

// scan arithmetic operation
func (l *Lexer) scanOperation() {
	op := l.expression[l.position]

	switch op {
	case "plus":
		l.tokens = append(l.tokens, Token{_type: ADD, Value: op})

	case "minus":
		l.tokens = append(l.tokens, Token{_type: SUB, Value: op})

	case "times":
		l.tokens = append(l.tokens, Token{_type: MUL, Value: op})

	case "divided by":
		l.tokens = append(l.tokens, Token{_type: DIV, Value: op})

	default:
		panic("unimplemented operation")
	}

}

func (l *Lexer) isOperation() bool {
	for _, l := range l.peek() {
		if !unicode.IsLetter(l) {
			return false
		}
	}

	if _, ok := valid_operations[l.peek()]; ok {
		return true
	}

	return false
}

func (l *Lexer) next() {
	l.position++
}

func (l *Lexer) peek() string {
	if l.position >= len(l.expression) {
		return ""
	}
	return l.expression[l.position]
}
