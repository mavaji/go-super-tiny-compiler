package compiler

import "fmt"

type Traverser struct {
}

func (t *Traverser) Traverse(ast *Node) error {
	err := t.traverseNode(ast, nil)
	if err != nil {
		return err
	}

	return nil
}

func (t *Traverser) traverseArray(array *[]*Node, parent *Node) error {
	for _, child := range *array {
		err := t.traverseNode(child, parent)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Traverser) traverseNode(node *Node, parent *Node) error {
	visitor := getVisitor(node.Type)
	if visitor != nil {
		visitor.Enter(node, parent)
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
		visitor.Exit(node, parent)
	}

	return nil
}

type Visitor interface {
	Enter(node *Node, parent *Node)
	Exit(node *Node, parent *Node)
}

type NumberLiteralVisitor struct {
}

func (n NumberLiteralVisitor) Enter(node *Node, parent *Node) {
	*parent.Context = append(*parent.Context, &Node{
		Type:  "NumberLiteral",
		Value: node.Value,
	})
}

func (n NumberLiteralVisitor) Exit(node *Node, parent *Node) {
}

type CallExpressionVisitor struct {
}

func (c CallExpressionVisitor) Enter(node *Node, parent *Node) {
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

func (c CallExpressionVisitor) Exit(node *Node, parent *Node) {
}

func getVisitor(nodeType string) Visitor {
	switch nodeType {
	case "NumberLiteral":
		return NumberLiteralVisitor{}
	case "CallExpression":
		return CallExpressionVisitor{}
	}
	return nil
}
