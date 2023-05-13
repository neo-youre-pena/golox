package interpreter

import (
	"github.com/neo-youre-pena/golox/src/token"
)

type Expr interface {
}

type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}
type Grouping struct {
	Expression Expr
}
type Literal struct {
	Value interface{}
}
type Unary struct {
	Operator token.Token
	Right    Expr
}
