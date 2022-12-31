package compiler

import (
	"go-super-tiny-compiler/src/compiler"
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("can build an AST from tokens", func(t *testing.T) {
		tokens := []compiler.Token{
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

		parser := compiler.Parser{}
		ast, err := parser.Parse(tokens)
		want := &compiler.Node{
			Type: "Program",
			Params: &[]*compiler.Node{
				{
					Type: "CallExpression",
					Name: "add",
					Params: &[]*compiler.Node{
						{
							Type:  "NumberLiteral",
							Value: "2",
						},
						{
							Type: "CallExpression",
							Name: "subtract",
							Params: &[]*compiler.Node{
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
		tokens := []compiler.Token{
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

		parser := compiler.Parser{}
		ast, err := parser.Parse(tokens)
		wantError := "syntax error: ')' at position: 0"
		if err.Error() != wantError {
			t.Errorf("want error %s but got %v", wantError, err)
		}
		if ast != nil {
			t.Errorf("want nil but got %v", ast)
		}
	})
}
