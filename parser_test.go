package main

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("can build an AST from tokens", func(t *testing.T) {
		tokens := []Token{
			{
				Type:  "paren",
				Value: "(",
			},
			{
				Type:  "name",
				Value: "add",
			},
			{
				Type:  "number",
				Value: "2",
			},
			{
				Type:  "paren",
				Value: "(",
			},
			{
				Type:  "name",
				Value: "subtract",
			},
			{
				Type:  "number",
				Value: "4",
			},
			{
				Type:  "number",
				Value: "2",
			},
			{
				Type:  "paren",
				Value: ")",
			},
			{
				Type:  "paren",
				Value: ")",
			},
		}

		parser := Parser{}
		ast, err := parser.parse(tokens)
		want := &Node{
			Type: "Program",
			Params: &[]*Node{
				{
					Type: "CallExpression",
					Name: "add",
					Params: &[]*Node{
						{
							Type:  "NumberLiteral",
							Value: "2",
						},
						{
							Type: "CallExpression",
							Name: "subtract",
							Params: &[]*Node{
								{
									Type:  "NumberLiteral",
									Value: "4",
								},
								{
									Type:  "NumberLiteral",
									Value: "2",
								},
							},
						},
					},
				},
			},
		}

		if err != nil {
			t.Errorf("want no error but got error: %s", err)
		}
		if !reflect.DeepEqual(ast, want) {
			t.Errorf("want %v but got %v", want, ast)
		}
	})

	t.Run("returns error for syntax errors", func(t *testing.T) {
		tokens := []Token{
			{
				Type:  "paren",
				Value: ")",
			},
			{
				Type:  "name",
				Value: "add",
			},
			{
				Type:  "number",
				Value: "2",
			},
			{
				Type:  "paren",
				Value: "(",
			},
			{
				Type:  "name",
				Value: "subtract",
			},
			{
				Type:  "number",
				Value: "4",
			},
			{
				Type:  "number",
				Value: "2",
			},
			{
				Type:  "paren",
				Value: ")",
			},
			{
				Type:  "paren",
				Value: ")",
			},
		}

		parser := Parser{}
		ast, err := parser.parse(tokens)
		wantError := "syntax error: ')' at position: 0"
		if err.Error() != wantError {
			t.Errorf("want error %s but got %v", wantError, err)
		}
		if ast != nil {
			t.Errorf("want nil but got %v", ast)
		}
	})
}
