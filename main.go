package main

import (
	"fmt"
	"strconv"
	"strings"
)

// reg := regexp.MustCompile(`[a-zA-Z0-9\s()+/*-]+`)

/*
var digits [20]string = [20]string{
"zero", "one", "two", "three",
"four", "five", "six", "seven",
"eight", "nine", "ten", "eleven",
"twelve", "thirteen", "fourteen", "fifteen",
"sixteen", "seventeen", "eighteen", "nineteen",
}


var composite_numbers [9]string = [9]string{
"twenty", "thirty", "forty",
"fifty", "sixty", "seventy",
"eigthy", "ninety", "one hundred",
}
*/

func main() {
	//message := flag.String("m", "", "message to parse")
	//flag.Parse()

	// TODO: make language sanitizer
	message := "one million twenty three"

	lexer := NewLexer(strings.Split(message, " "))
	tokens := lexer.Scan()
	fmt.Println(tokens)

	parser := NewParser(tokens)

	result := parser.Parse()

	fmt.Println(result)
}

func wordToNumber(words []string) ([]string, error) {
	expression := []string{}

	for i, w := range words {
		num, ok := digits[w]
		if !ok {
			op, ok := valid_operations[w]
			if !ok {
				return nil, fmt.Errorf("unknown word: %s", w)
			}

			expression = append(expression, string(op))
			continue
		}

		if len(words) > i+2 && words[i+1] == "and" {

		}

		expression = append(expression, strconv.Itoa(num))

	}

	return expression, nil
}
