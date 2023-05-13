package interpreter

import (
	"lox/token"
	"strings"
)

type Visitor interface {
	VisitBinary(expr *Binary) any
	VisitGrouping(expr *Grouping) any
	VisitLiteral(expr *Literal) any
	VisitUnary(expr *Unary) any
}

type Expr interface {
	String() string
	Accept(visitor Visitor) any
}

type Binary struct {
	Expr
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (b *Binary) String() string {
	return Parenthesize(b.Operator.Lexeme, b.Left, b.Right)
}

func (b *Binary) Accept(visitor Visitor) any {
	return visitor.VisitBinary(b)
}

type Grouping struct {
	Expr
	Expression Expr
}

func (g *Grouping) String() string {
	return Parenthesize("group", g.Expression)
}

type Literal struct {
	Expr
	Value interface{}
}

func (l *Literal) String() string {
	if l.Value == nil {
		return "nil"
	}
	return l.Value.(string)
}

type Unary struct {
	Expr
	Operator token.Token
	Right    Expr
}

func (u *Unary) String() string {
	return Parenthesize(u, u.Operator.Lexeme, u.Right)
}

func Parenthesize(ee Expr, name string, exprs ...Expr) string {
	var sb strings.Builder
	sb.WriteString("(")
	sb.WriteString(name)
	for _, e := range exprs {
		sb.WriteString(" ")
		sb.WriteString(e.Accept(ee))
	}
	sb.WriteString(")")
	return sb.String()
}
