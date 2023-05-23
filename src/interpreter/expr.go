package interpreter

import (
	"github.com/neo-youre-pena/golox/src/token"
)

type Expr[R any] interface {
	accept(v Visitor[R]) R
}

type Visitor[R any] interface {
	visitForBinary(*Binary[R]) R
	visitForGrouping(*Grouping[R]) R
	visitForLiteral(*Literal[R]) R
	visitForUnary(*Unary[R]) R
}
type Binary[R any] struct {
	Left     Expr[R]
	Operator token.Token
	Right    Expr[R]
}

func (c *Binary[R]) accept(v Visitor[R]) R {
	return v.visitForBinary(c)
}

type Grouping[R any] struct {
	Expression Expr[R]
}

func (c *Grouping[R]) accept(v Visitor[R]) R {
	return v.visitForGrouping(c)
}

type Literal[R any] struct {
	Value interface{}
}

func (c *Literal[R]) accept(v Visitor[R]) R {
	return v.visitForLiteral(c)
}

type Unary[R any] struct {
	Operator token.Token
	Right    Expr[R]
}

func (c *Unary[R]) accept(v Visitor[R]) R {
	return v.visitForUnary(c)
}
