package ast

import (
	"slang/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// The root of every AST 
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// bind a value to a given name.
// var x = 5 -> name = x and value = 5
type VarStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

func (vs *VarStatement) statementNode() {

}

func (vs *VarStatement) TokenLiteral() string {
	return vs.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}