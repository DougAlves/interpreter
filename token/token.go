package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSING    = "="
	PLUS      = "+"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	FUNCION   = "FUNCTION"
	LET       = "LET"
	RETURN    = "RETURN"
	IF        = "IF"
	ELSE      = "ELSE"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	BANG      = "!"
	MINUS     = "-"
	SLASH     = "/"
	ASTERISK  = "*"
	LT        = "<"
	GT        = ">"
	EQ        = "=="
	NEQ       = "!="
)

var keywords = map[string]TokenType{
	"fn":     FUNCION,
	"let":    LET,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
