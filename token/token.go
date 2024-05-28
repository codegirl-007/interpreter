package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// Different types of token types
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y
	INT   = "INT"   // 123456

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET" // I'm going to change this to var at some point
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// checks the keywords table to see whether the given identifier is in fact a keyword
// If it is a keyword, return the corresponding token type
// If it isn't, we get back Token.IDENT for user-defined identifiers
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
