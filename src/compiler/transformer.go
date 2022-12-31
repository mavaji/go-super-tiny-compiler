package compiler

type Transformer struct {
}

func (t *Transformer) Transform(ast *Node) (*Node, error) {
	newAst := &Node{
		Type:   "Program",
		Params: &[]*Node{},
	}

	ast.Context = newAst.Params

	traverser := Traverser{}
	err := traverser.Traverse(ast)
	if err != nil {
		return nil, err
	}

	return newAst, nil
}
