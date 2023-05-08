package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	elements   []Token
	position   int
	expression strings.Builder
}

func NewParser(elem []Token) *Parser {
	return &Parser{
		elements:   elem,
		position:   0,
		expression: strings.Builder{},
	}
}

func (p *Parser) Parse() string {
	for p.position < len(p.elements) {
		switch p.peek()._type {
		case INT:
			p.expression.WriteString(p.parseToken())
			p.position++

		default:
			break
		}
	}

	return p.expression.String()
}

func (p *Parser) parseToken() string {
	actualPosition := p.position
	result := 0
	num := 0

	for actualPosition+1 < len(p.elements) {
		if p.elements[actualPosition+1]._type != INT {
			break
		}

		num, _ = digits[p.elements[actualPosition].Value]
		nextNum, _ := digits[p.elements[actualPosition+1].Value]

	Cond:
		switch {
		case isSingleDigit(p.elements[actualPosition].Value) && isCentesimalOrAbove(p.elements[actualPosition+1].Value):
			result += num * nextNum
			break Cond
		case isCentesimalOrAbove(p.elements[actualPosition].Value) && !isSingleDigit(p.elements[actualPosition+1].Value):
			result += nextNum
			break Cond
		case (!isSingleDigit(p.elements[actualPosition].Value) && !isCentesimalOrAbove(p.elements[actualPosition].Value)) && isSingleDigit(p.elements[actualPosition+1].Value):
			result += nextNum
			break Cond
		default:
			panic("invalid token sequence in number parsing")
		}

		fmt.Printf("primeiro: %d\n", num)
		fmt.Printf("proximo: %d\n", nextNum)
		fmt.Printf("resultado: %d\n\n", result)
		actualPosition++
	}

	p.position = actualPosition

	return strconv.Itoa(result)
}

func (p *Parser) peek() *Token {
	if p.position >= len(p.elements) {
		return nil
	}

	return &(p.elements[p.position])
}
