package main

import (
	"reflect"
	"testing"
)

func TestTransformer(t *testing.T) {
	t.Run("can transform AST to a new AST", func(t *testing.T) {
		transformer := Transformer{}
		ast := &Node{
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
		newAst, err := transformer.transform(ast)
		want := &Node{
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

		if err != nil {
			t.Errorf("want no error but got error: %s", err)
		}
		if !reflect.DeepEqual(newAst, want) {
			t.Errorf("want %v but got %v", want, newAst)
		}
	})

	t.Run("returns error when unknown type happens", func(t *testing.T) {
		transformer := Transformer{}
		ast := &Node{
			Type: "Program",
			Params: &[]*Node{
				{
					Type: "CallExpression",
					Name: "add",
					Params: &[]*Node{
						{
							Type:  "NumberLiteralLPJLJL",
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
		newAst, err := transformer.transform(ast)
		wantError := "unknown type error: NumberLiteralLPJLJL"
		if err.Error() != wantError {
			t.Errorf("want error %s but got %v", wantError, err)
		}
		if newAst != nil {
			t.Errorf("want nil but got %v", newAst)
		}
	})
}
