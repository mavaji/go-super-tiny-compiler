package compiler

import (
	"fmt"
	"regexp"
)

type tokenizer struct {
	whiteSpace *regexp.Regexp
	numbers    *regexp.Regexp
	letters    *regexp.Regexp
}

func NewTokenizer() tokenizer {
	return tokenizer{
		whiteSpace: regexp.MustCompile(`\s`),
		numbers:    regexp.MustCompile(`[0-9]`),
		letters:    regexp.MustCompile(`[a-z]`),
	}
}

func (t *tokenizer) Tokenize(input string) ([]Token, error) {
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

		if t.whiteSpace.MatchString(string(char)) {
			current += 1
			continue
		}

		if t.numbers.MatchString(string(char)) {
			value := ""
			for t.numbers.MatchString(string(char)) {
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

		if t.letters.MatchString(string(char)) {
			value := ""
			for t.letters.MatchString(string(char)) {
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
