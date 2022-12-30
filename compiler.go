package main

type Compiler struct {
}

func (c *Compiler) compile(input string) (string, error) {
	tokenizer := Tokenizer{}
	parser := Parser{}
	transformer := Transformer{}
	codeGenerator := CodeGenerator{}

	tokens, err := tokenizer.Tokenize(input)
	if err != nil {
		return "", err
	}

	ast, err := parser.parse(tokens)
	if err != nil {
		return "", err
	}

	newAst, err := transformer.transform(ast)
	if err != nil {
		return "", err
	}

	output, err := codeGenerator.generateCode(newAst)
	if err != nil {
		return "", err
	}

	return output, nil
}
