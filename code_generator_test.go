package main

import (
	"testing"
)

func TestCodeGenerator(t *testing.T) {
	t.Run("can generate code from AST", func(t *testing.T) {
		codeGenerator := CodeGenerator{}
		ast := &Node{
			Type: "Program",
			Params: &[]*Node{
				{
					Type: "ExpressionStatement",
					Expression: &Node{
						Type: "CallExpression",
						Callee: &Node{
							Type: "Identifier",
							Name: "add",
						},
						Arguments: &[]*Node{
							{
								Type:  "NumberLiteral",
								Value: "2",
							},
							{
								Type: "CallExpression",
								Callee: &Node{
									Type: "Identifier",
									Name: "subtract",
								},
								Arguments: &[]*Node{
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
			},
		}
		code, err := codeGenerator.generateCode(ast)
		want := "add(2, subtract(4, 2));"
		if err != nil {
			t.Errorf("want no error but got error: %s", err)
		}
		if code != want {
			t.Errorf("want %v but got %v", want, code)
		}
	})

	t.Run("returns error when unknown node type happens", func(t *testing.T) {
		codeGenerator := CodeGenerator{}
		ast := &Node{
			Type: "Program",
			Params: &[]*Node{
				{
					Type: "ExpressionStatement",
					Expression: &Node{
						Type: "CallExpression",
						Callee: &Node{
							Type: "Identifier",
							Name: "add",
						},
						Arguments: &[]*Node{
							{
								Type:  "NumberLiteralLPJLJL",
								Value: "2",
							},
							{
								Type: "CallExpression",
								Callee: &Node{
									Type: "Identifier",
									Name: "subtract",
								},
								Arguments: &[]*Node{
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
			},
		}
		code, err := codeGenerator.generateCode(ast)
		wantError := "unknown type error: NumberLiteralLPJLJL"
		if err.Error() != wantError {
			t.Errorf("want error %s but got %v", wantError, err)
		}
		if code != "" {
			t.Errorf("want empty string but got %v", code)
		}
	})
}
