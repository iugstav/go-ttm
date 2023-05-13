package main

import (
	"fmt"
	"strings"
)

// the unique usage of this function is to fix
// the [..., "divided", "by", ...] to [..., "divided by", ...]
func fixInputOperations(input *[]string) {
	for i := 0; i < len(*input); i++ {
		if (*input)[i] != "divided" {
			continue
		}

		(*input)[i] = "divided by"
		*input = append((*input)[:i+1], (*input)[i+2:]...)
	}
}

func main() {
	// message := flag.String("m", "", "message to parse")
	// flag.Parse()

	message := "two million three thousand twenty three plus two times twenty nine divided by twenty four times twenty"
	input := strings.Split(message, " ")
	fixInputOperations(&input)

	lexer := NewLexer(input)
	tokens := lexer.Scan()
	fmt.Println(tokens)

	parser := NewParser(tokens)

	result := parser.Parse()

	fmt.Println(result)
}
