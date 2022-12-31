package compiler

import (
	"fmt"
)

type parser struct {
}

func NewParser() parser {
	return parser{}
}

func (p *parser) Parse(tokens []Token) (*Node, error) {
	current := 0
	var walk func() (*Node, error)
	walk = func() (*Node, error) {
		token := tokens[current]
		if token.Type == "number" {
			current += 1
			return &Node{
				Type:  "NumberLiteral",
				Value: token.Value,
			}, nil
		}

		if token.Type == "string" {
			current += 1
			return &Node{
				Type:  "StringLiteral",
				Value: token.Value,
			}, nil
		}

		if token.Type == "paren" && token.Value == "(" {
			current += 1
			token = tokens[current]

			node := &Node{
				Type:   "CallExpression",
				Name:   token.Value,
				Params: &[]*Node{},
			}

			current += 1
			token = tokens[current]

			for (token.Type != "paren") || (token.Type == "paren" && token.Value != ")") {
				node1, err := walk()
				if err != nil {
					return nil, err
				}
				*node.Params = append(*node.Params, node1)
				token = tokens[current]
			}

			current += 1
			return node, nil
		}

		return nil, fmt.Errorf("syntax error: '%s' at position: %d", token.Value, current)
	}

	ast := &Node{
		Type:   "Program",
		Params: &[]*Node{},
	}

	for current < len(tokens) {
		node, err := walk()
		if err != nil {
			return nil, err
		}
		*ast.Params = append(*ast.Params, node)
	}

	return ast, nil
}
