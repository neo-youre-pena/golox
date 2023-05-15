package interpreter

import (
	"github.com/neo-youre-pena/golox/src/token"
)

type Expr interface {
	accept(v Visitor) string
}

type Visitor interface {
	visitForBinary(*Binary) string
	visitForGrouping(*Grouping) string
	visitForLiteral(*Literal) string
	visitForUnary(*Unary) string
}
type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (c *Binary) accept(v Visitor) string {
	return v.visitForBinary(c)
}

type Grouping struct {
	Expression Expr
}

func (c *Grouping) accept(v Visitor) string {
	return v.visitForGrouping(c)
}

type Literal struct {
	Value interface{}
}

func (c *Literal) accept(v Visitor) string {
	return v.visitForLiteral(c)
}

type Unary struct {
	Operator token.Token
	Right    Expr
}

func (c *Unary) accept(v Visitor) string {
	return v.visitForUnary(c)
}
