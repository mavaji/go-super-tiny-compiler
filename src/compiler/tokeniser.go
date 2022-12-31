package compiler

import (
	"fmt"
	"regexp"
)

type Tokenizer struct {
}

func (t *Tokenizer) Tokenize(input string) ([]Token, error) {
	current := 0
	var tokens []Token
	for current < len(input) {
		char := input[current]

		if char == '(' {
			tokens = append(tokens, Token{
				Type:  "paren",
				Value: "(",
			})
			current += 1
			continue
		}

		if char == ')' {
			tokens = append(tokens, Token{
				Type:  "paren",
				Value: ")",
			})
			current += 1
			continue
		}

		whiteSpace := regexp.MustCompile(`\s`)
		if whiteSpace.MatchString(string(char)) {
			current += 1
			continue
		}

		numbers := regexp.MustCompile(`[0-9]`)
		if numbers.MatchString(string(char)) {
			value := ""
			for numbers.MatchString(string(char)) {
				value += string(char)
				current += 1
				char = input[current]
			}

			tokens = append(tokens, Token{
				Type:  "number",
				Value: value,
			})
			continue
		}

		letters := regexp.MustCompile(`[a-z]`)
		if letters.MatchString(string(char)) {
			value := ""
			for letters.MatchString(string(char)) {
				value += string(char)
				current += 1
				char = input[current]
			}

			tokens = append(tokens, Token{
				Type:  "name",
				Value: value,
			})
			continue
		}

		return nil, fmt.Errorf("lexical error: '%c' at position: %d", char, current)
	}

	return tokens, nil
}
