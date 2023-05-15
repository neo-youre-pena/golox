package interpreter

import (
	"testing"

	"github.com/neo-youre-pena/golox/src/token"
)

func TestAst(t *testing.T) {
	expr := Binary{
		Left: &Unary{
			Operator: token.Token{
				Type:    token.MINUS,
				Lexeme:  "-",
				Literal: nil,
				Line:    1,
			},
			Right: &Literal{
				Value: 123,
			},
		},
		Operator: token.Token{
			Type:    token.STAR,
			Lexeme:  "*",
			Literal: nil,
			Line:    1,
		},
		Right: &Grouping{
			Expression: &Literal{
				Value: 45.67,
			},
		},
	}

	ast := Ast{}
	t.Log(ast.print(&expr))
}
