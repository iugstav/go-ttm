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

// splits the slice of numbers of a token into pairs,
// so it represents a whole number
func chunkTokenValue(value []string) (pairs [][]string) {
	end := 0
	valueLength := len(value)

	for i := 0; i < valueLength; i += 2 {
		end = i + 2

		if end > valueLength {
			end = valueLength
		}

		pairs = append(pairs, value[i:end])
	}

	return pairs
}

func (p *Parser) Parse() string {
	for p.position < len(p.elements) {
		switch p.peek()._type {
		case INT:
			p.expression.WriteString(p.parseNumber())
			p.next()

		case ADD:
			p.expression.WriteString(" + ")
			p.next()

		case SUB:
			p.expression.WriteString(" - ")
			p.next()
		case MUL:
			p.expression.WriteString(" * ")
			p.next()

		case DIV:
			p.expression.WriteString(" / ")
			p.next()

		default:
			panic("illegal token in parsing")
		}
	}

	return p.expression.String()
}

func (p *Parser) parseNumber() string {
	result := 0
	num := 0
	tokenValue := p.peek().Value

	if strings.Index(tokenValue, " ") == -1 {
		num, _ = digits[tokenValue]
		return strconv.Itoa(num)
	}

	pairs := chunkTokenValue(strings.Split(tokenValue, " "))
	for i := 0; i < len(pairs); i++ {
		if len(pairs[i]) == 1 {
			num, _ = digits[pairs[i][0]]
			result += num

			break
		}
		// PAREI AQUI, FALTA RODAR O CODIGO E VER SE O SWITCH BATE
		num := digits[pairs[i][0]]
		nextNum := digits[pairs[i][1]]

		switch {
		case isSingleDigit(pairs[i][0]) && isCentesimalOrAbove(pairs[i][1]):
			result += num * nextNum
		case isCentesimalOrAbove(pairs[i][0]) && !isSingleDigit(pairs[i][1]):
			result += nextNum
		case (!isSingleDigit(pairs[i][0]) && !isCentesimalOrAbove(pairs[i][0])) && isSingleDigit(pairs[i][1]):
			result += num + nextNum
		default:
			panic(fmt.Sprintf("invalid token sequence in number parsing: %d at %d", num, i))
		}
	}
	//
	// for actualPosition+1 < len(pairs[i][0]) {
	// 	if pairs[i][0][actualPosition+1]._type != INT {
	// 		break
	// 	}
	//
	// 	num, _ = digits[pairs[i][0][actualPosition].Value]
	// 	nextNum, _ := digits[pairs[i][0][actualPosition+1].Value]
	//
	// Cond:
	// 	switch {
	// 	case isSingleDigit(pairs[i][0][actualPosition].Value) && isCentesimalOrAbove(p.elements[actualPosition+1].Value):
	// 		result += num * nextNum
	// 		break Cond
	// 	case isCentesimalOrAbove(pairs[i][0][actualPosition].Value) && !isSingleDigit(p.elements[actualPosition+1].Value):
	// 		result += nextNum
	// 		break Cond
	// 	case (!isSingleDigit(pairs[i][0][actualPosition].Value) && !isCentesimalOrAbove(p.elements[actualPosition].Value)) && isSingleDigit(p.elements[actualPosition+1].Value):
	// 		result += nextNum
	// 		fmt.Println(num)
	// 		fmt.Println(nextNum)
	// 		break Cond
	// 	default:
	// 		panic(fmt.Sprintf("invalid token sequence in number parsing: %d at %d", num, actualPosition))
	// 	}
	//
	// 	actualPosition++
	// }
	//
	// p.position = actualPosition

	return strconv.Itoa(result)
}

func (p *Parser) peek() *Token {
	if p.position >= len(p.elements) {
		return nil
	}

	return &(p.elements[p.position])
}

func (p *Parser) next() {
	p.position++
}
