package interpreter

import (
	"testing"

	"github.com/neo-youre-pena/golox/src/token"
)

func TestAst(t *testing.T) {
	expr := Binary[string]{
		Left: &Unary[string]{
			Operator: token.Token{
				Type:    token.MINUS,
				Lexeme:  "-",
				Literal: nil,
				Line:    1,
			},
			Right: &Literal[string]{
				Value: 123,
			},
		},
		Operator: token.Token{
			Type:    token.STAR,
			Lexeme:  "*",
			Literal: nil,
			Line:    1,
		},
		Right: &Grouping[string]{
			Expression: &Literal[string]{
				Value: 45.67,
			},
		},
	}

	ast := Ast{}
	t.Log(ast.print(&expr))
}
