package compiler

type Token struct {
	Type  string
	Value string
}

type Node struct {
	Type       string
	Name       string
	Value      string
	Callee     *Node
	Expression *Node
	Params     *[]*Node
	Arguments  *[]*Node
	Context    *[]*Node
}
