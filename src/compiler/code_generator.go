package compiler

import (
	"fmt"
	"strings"
)

type codeGenerator struct {
}

func NewCodeGenerator() codeGenerator {
	return codeGenerator{}
}

func (g *codeGenerator) GenerateCode(node *Node) (string, error) {
	switch node.Type {
	case "Program":
		return g.programCode(node)
	case "ExpressionStatement":
		return g.expressionStatementCode(node)
	case "CallExpression":
		return g.callExpressionCode(node)
	case "Identifier":
		return g.identifierCode(node)
	case "NumberLiteral":
		return g.numberLiteralCode(node)
	default:
		return "", fmt.Errorf("unknown type error: %s", node.Type)
	}
}

func (g *codeGenerator) numberLiteralCode(node *Node) (string, error) {
	return node.Value, nil
}

func (g *codeGenerator) identifierCode(node *Node) (string, error) {
	return node.Name, nil
}

func (g *codeGenerator) callExpressionCode(node *Node) (string, error) {
	calleeCode, err := g.GenerateCode(node.Callee)
	if err != nil {
		return "", err
	}

	var argumentsCode []string
	for _, argument := range *node.Arguments {
		argumentCode, err := g.GenerateCode(argument)
		if err != nil {
			return "", err
		}
		argumentsCode = append(argumentsCode, argumentCode)
	}

	return calleeCode + "(" + strings.Join(argumentsCode, ", ") + ")", nil
}

func (g *codeGenerator) expressionStatementCode(node *Node) (string, error) {
	expressionCode, err := g.GenerateCode(node.Expression)
	if err != nil {
		return "", err
	}

	return expressionCode + ";", nil
}

func (g *codeGenerator) programCode(node *Node) (string, error) {
	var programCode []string

	for _, param := range *node.Params {
		code, err := g.GenerateCode(param)
		if err != nil {
			return "", err
		}
		programCode = append(programCode, code)
	}

	return strings.Join(programCode, "\n"), nil
}
