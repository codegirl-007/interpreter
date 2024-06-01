package parser

import (
	"fmt"
	"slang/ast"
	"slang/lexer"
	"slang/token"
)

type Parser struct {
	l *lexer.Lexer

	currToken token.Token // currToken and peekToken are like the pointers in that other file
	errors    []string
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}
func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.VAR:
		return p.parseVarStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseVarStatement() *ast.VarStatement {
	// construct ast node with the node currently on
	stmt := &ast.VarStatement{Token: p.currToken}

	// check to see what the next token is (because see you're using expectPeek here which advances the pointer)
	if !p.expectPeek(token.IDENT) { // expect an IDENT
		// returning nil results in the entire statement being ignored because of an error in input
		// peekError is called in expect token so you will see errors.
		return nil
	}

	// create an ast node for the identifier
	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	// expect an equal sign
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: jump over until semicolon -- revisit this
	for !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currToken}

	p.nextToken()

	for !p.currTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) currTokenIs(t token.TokenType) bool {
	return p.currToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// enforce correctness of the order of tokens by checking the type
// of the next token. Only advance pointers if it is the expected token
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) ParseProgram() *ast.Program {
	// construct the root node of the AST
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// Iterate over every token until EOF is encountered
	for p.currToken.Type != token.EOF {
		stmt := p.parseStatement() // parse the statements
		if stmt != nil {
			program.Statements = append(program.Statements, stmt) // and then add them to the AST
		}
		p.nextToken() // advance current token and peek token
	}

	return program
}
