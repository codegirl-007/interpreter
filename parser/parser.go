package parser

import (
	"slang/ast"
	"slang/lexer"
	"slang/token"
)

type Parser struct {
	l *lexer.Lexer

	currToken token.Token // currToken and peekToken are like the pointers in that other file
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
