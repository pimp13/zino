package ast

type Node interface {
	String() string
}

type Expression interface {
	Node
	expressionNode()
}

type NumberLiteral struct {
	Value string
}

func (n *NumberLiteral) expressionNode() {}
func (n *NumberLiteral) String() string  { return n.Value }

type Identifier struct {
	Name string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) String() string  { return i.Name }

type InfixExpression struct {
	Left     Expression
	Operator string
	Right    Expression
}

func (i *InfixExpression) expressionNode() {}
func (i *InfixExpression) String() string {
	return "(" + i.Left.String() + " " + i.Operator + " " + i.Right.String() + ")"
}
