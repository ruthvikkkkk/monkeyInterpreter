package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"for":    FOR,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	ILLEGAl = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers + Literals
	IDENT = "IDENT"
	INT   = "INT"

	//Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	PERCENT  = "%"
	CARET    = "^"

	BANG   = "!"
	EQ     = "=="
	NOT_EQ = "!="

	LT       = "<"
	LT_OR_EQ = "<="

	GT       = ">"
	GT_OR_EQ = ">="

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "true"
	FALSE    = "false"
	IF       = "IF"
	ELSE     = "ELSE"
	FOR      = "FOR"
	RETURN   = "return"
)
