package compiler

import "fmt"

type Transformer struct {
}

func (t *Transformer) Transform(ast *Node) (*Node, error) {
	newAst := &Node{
		Type:   "Program",
		Params: &[]*Node{},
	}

	ast.Context = newAst.Params

	err := t.traverse(ast)
	if err != nil {
		return nil, err
	}

	return newAst, nil
}

func (t *Transformer) traverse(ast *Node) error {
	err := t.traverseNode(ast, nil)
	if err != nil {
		return err
	}

	return nil
}

func (t *Transformer) traverseArray(array *[]*Node, parent *Node) error {
	for _, child := range *array {
		err := t.traverseNode(child, parent)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Transformer) traverseNode(node *Node, parent *Node) error {
	visitor := t.getVisitor(node.Type)
	if visitor != nil {
		visitor.enter(node, parent)
	}

	switch node.Type {
	case "Program":
		err := t.traverseArray(node.Params, node)
		if err != nil {
			return err
		}
	case "CallExpression":
		err := t.traverseArray(node.Params, node)
		if err != nil {
			return err
		}
	case "NumberLiteral":
		break
	default:
		return fmt.Errorf("unknown type error: %s", node.Type)
	}

	if visitor != nil {
		visitor.exit(node, parent)
	}

	return nil
}

type visitor interface {
	enter(node *Node, parent *Node)
	exit(node *Node, parent *Node)
}

type numberLiteralVisitor struct {
}

func (n numberLiteralVisitor) enter(node *Node, parent *Node) {
	*parent.Context = append(*parent.Context, &Node{
		Type:  "NumberLiteral",
		Value: node.Value,
	})
}

func (n numberLiteralVisitor) exit(node *Node, parent *Node) {
}

type callExpressionVisitor struct {
}

func (c callExpressionVisitor) enter(node *Node, parent *Node) {
	expression := &Node{
		Type: "CallExpression",
		Callee: &Node{
			Type: "Identifier",
			Name: node.Name,
		},
		Arguments: &[]*Node{},
	}

	node.Context = expression.Arguments

	if parent.Type != "CallExpression" {
		expression = &Node{
			Type:       "ExpressionStatement",
			Expression: expression,
		}
	}

	*parent.Context = append(*parent.Context, expression)
}

func (c callExpressionVisitor) exit(node *Node, parent *Node) {
}

func (t *Transformer) getVisitor(nodeType string) visitor {
	switch nodeType {
	case "NumberLiteral":
		return numberLiteralVisitor{}
	case "CallExpression":
		return callExpressionVisitor{}
	}
	return nil
}
