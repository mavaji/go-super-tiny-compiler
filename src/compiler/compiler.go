package compiler

type Compiler struct {
}

func (c *Compiler) Compile(input string) (string, error) {
	tokenizer := Tokenizer{}
	parser := Parser{}
	transformer := Transformer{}
	codeGenerator := CodeGenerator{}

	tokens, err := tokenizer.Tokenize(input)
	if err != nil {
		return "", err
	}

	ast, err := parser.Parse(tokens)
	if err != nil {
		return "", err
	}

	newAst, err := transformer.Transform(ast)
	if err != nil {
		return "", err
	}

	output, err := codeGenerator.GenerateCode(newAst)
	if err != nil {
		return "", err
	}

	return output, nil
}
