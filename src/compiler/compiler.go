package compiler

type compiler struct {
	tokenizer     tokenizer
	parser        parser
	transformer   transformer
	codeGenerator codeGenerator
}

func NewCompiler() compiler {
	return compiler{
		tokenizer:     NewTokenizer(),
		parser:        NewParser(),
		transformer:   NewTransformer(),
		codeGenerator: NewCodeGenerator(),
	}
}
func (c *compiler) Compile(input string) (string, error) {
	tokens, err := c.tokenizer.Tokenize(input)
	if err != nil {
		return "", err
	}

	ast, err := c.parser.Parse(tokens)
	if err != nil {
		return "", err
	}

	newAst, err := c.transformer.Transform(ast)
	if err != nil {
		return "", err
	}

	output, err := c.codeGenerator.GenerateCode(newAst)
	if err != nil {
		return "", err
	}

	return output, nil
}
