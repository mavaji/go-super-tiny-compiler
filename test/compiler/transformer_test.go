package compiler

import (
	"go-super-tiny-compiler/src/compiler"
	"reflect"
	"testing"
)

func TestTransformer(t *testing.T) {
	t.Run("can transform AST to a new AST", func(t *testing.T) {
		transformer := compiler.Transformer{}
		ast := &compiler.Node{
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
		newAst, err := transformer.Transform(ast)
		want := &compiler.Node{
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

		if err != nil {
			t.Errorf("want no error but got error: %s", err)
		}
		if !reflect.DeepEqual(newAst, want) {
			t.Errorf("want %v but got %v", want, newAst)
		}
	})

	t.Run("returns error when unknown type happens", func(t *testing.T) {
		transformer := compiler.Transformer{}
		ast := &compiler.Node{
			Type: "Program",
			Params: &[]*compiler.Node{
				{
					Type: "CallExpression",
					Name: "add",
					Params: &[]*compiler.Node{
						{
							Type:  "NumberLiteralLPJLJL",
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
		newAst, err := transformer.Transform(ast)
		wantError := "unknown type error: NumberLiteralLPJLJL"
		if err.Error() != wantError {
			t.Errorf("want error %s but got %v", wantError, err)
		}
		if newAst != nil {
			t.Errorf("want nil but got %v", newAst)
		}
	})
}
