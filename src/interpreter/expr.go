package interpreter

import (
	"github.com/neo-youre-pena/golox/src/token"
)

type Expr interface {
}

type Visitor interface {
	visitForBinary(*Binary)
	visitForGrouping(*Grouping)
	visitForLiteral(*Literal)
	visitForUnary(*Unary)
}
type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (c *Binary) accept(v Visitor) {
	v.visitForBinary(c)
}

type Grouping struct {
	Expression Expr
}

func (c *Grouping) accept(v Visitor) {
	v.visitForGrouping(c)
}

type Literal struct {
	Value interface{}
}

func (c *Literal) accept(v Visitor) {
	v.visitForLiteral(c)
}

type Unary struct {
	Operator token.Token
	Right    Expr
}

func (c *Unary) accept(v Visitor) {
	v.visitForUnary(c)
}
