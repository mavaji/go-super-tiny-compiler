package compiler

import (
	"fmt"
	"strings"
)

type CodeGenerator struct {
}

func (g *CodeGenerator) GenerateCode(node *Node) (string, error) {
	switch node.Type {
	case "Program":
		var result []string
		for _, param := range *node.Params {
			code, err := g.GenerateCode(param)
			if err != nil {
				return "", err
			}
			result = append(result, code)
		}
		return strings.Join(result, "\n"), nil
	case "ExpressionStatement":
		code, err := g.GenerateCode(node.Expression)
		if err != nil {
			return "", err
		}
		return code + ";", nil
	case "CallExpression":
		code, err := g.GenerateCode(node.Callee)
		if err != nil {
			return "", err
		}
		var result []string
		for _, argument := range *node.Arguments {
			code, err := g.GenerateCode(argument)
			if err != nil {
				return "", err
			}
			result = append(result, code)
		}
		return code + "(" + strings.Join(result, ", ") + ")", nil
	case "Identifier":
		return node.Name, nil
	case "NumberLiteral":
		return node.Value, nil
	default:
		return "", fmt.Errorf("unknown type error: %s", node.Type)
	}
}
