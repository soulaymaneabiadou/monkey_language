package token

type TokenType string

type Token struct {
	Type    TokenType // the token type
	Literal string    // the value
}

// allowed TokenTypes
const (
	// special types
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// indentifiers and literals
	IDENT = "IDENT" // user-defined identifiers/variable names
	INT   = "INT"   // integers

	// operators for math and conditionals
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"
	BANG     = "!"

	LT = "<"
	GT = ">"

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	EQ     = "=="
	NOT_EQ = "!="

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// lookup an identifier and decide if it is a keyword or a user-defined one
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
