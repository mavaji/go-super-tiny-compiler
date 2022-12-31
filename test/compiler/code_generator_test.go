package compiler

import (
	"go-super-tiny-compiler/src/compiler"
	"testing"
)

func TestCodeGenerator(t *testing.T) {
	t.Run("can generate code from AST", func(t *testing.T) {
		codeGenerator := compiler.NewCodeGenerator()
		ast := &compiler.Node{
			Type: "Program",
			Params: &[]*compiler.Node{
				{
					Type: "ExpressionStatement",
					Expression: &compiler.Node{
						Type: "CallExpression",
						Callee: &compiler.Node{
							Type: "Identifier",
							Name: "add",
						},
						Arguments: &[]*compiler.Node{
							{
								Type:  "NumberLiteral",
								Value: "2",
							},
							{
								Type: "CallExpression",
								Callee: &compiler.Node{
									Type: "Identifier",
									Name: "subtract",
								},
								Arguments: &[]*compiler.Node{
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
		code, err := codeGenerator.GenerateCode(ast)
		want := "add(2, subtract(4, 2));"
		if err != nil {
			t.Errorf("want no error but got error: %s", err)
		}
		if code != want {
			t.Errorf("want %v but got %v", want, code)
		}
	})

	t.Run("returns error when unknown node type happens", func(t *testing.T) {
		codeGenerator := compiler.NewCodeGenerator()
		ast := &compiler.Node{
			Type: "Program",
			Params: &[]*compiler.Node{
				{
					Type: "ExpressionStatement",
					Expression: &compiler.Node{
						Type: "CallExpression",
						Callee: &compiler.Node{
							Type: "Identifier",
							Name: "add",
						},
						Arguments: &[]*compiler.Node{
							{
								Type:  "NumberLiteralLPJLJL",
								Value: "2",
							},
							{
								Type: "CallExpression",
								Callee: &compiler.Node{
									Type: "Identifier",
									Name: "subtract",
								},
								Arguments: &[]*compiler.Node{
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
		code, err := codeGenerator.GenerateCode(ast)
		wantError := "unknown type error: NumberLiteralLPJLJL"
		if err.Error() != wantError {
			t.Errorf("want error %s but got %v", wantError, err)
		}
		if code != "" {
			t.Errorf("want empty string but got %v", code)
		}
	})
}
