package interpreter

import (
	"bytes"
	"fmt"
)

type Ast struct {
}

func (a *Ast) print(expr Expr) string {
	return expr.accept(a)
}

func (a *Ast) visitForBinary(b *Binary) string {
	return a.parenthesize(b.Operator.Lexeme, b.Left, b.Right)
}

func (a *Ast) visitForGrouping(g *Grouping) string {
	return a.parenthesize("group", g.Expression)
}

func (a *Ast) visitForLiteral(l *Literal) string {
	if l.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", l.Value)
}

func (a *Ast) visitForUnary(u *Unary) string {
	return a.parenthesize(u.Operator.Lexeme, u.Right)
}

func (a *Ast) parenthesize(name string, exprs ...Expr) string {
	var buffer bytes.Buffer

	buffer.WriteString("(")
	buffer.WriteString(name)

	for _, expr := range exprs {
		buffer.WriteString(" ")
		buffer.WriteString(expr.accept(a))
	}

	buffer.WriteString(")")

	return buffer.String()
}
