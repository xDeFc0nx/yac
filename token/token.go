package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLIEGAL = "ILLIEGAL"
	EOF      = "EOF"
	// indendifiers & literals
	IDENT = "IDENT"
	INT   = "INT"
	// operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"
	// Delimiters

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNKSION"
	LET      = "LEJ"
	TRUE     = "PO"
	FALSE    = "JO"
	IF       = "NESE"
	ELSE     = "perndryshe"
	RETURN   = "KTHE"
)

var keywords = map[string]TokenType{
	"fn":         FUNCTION,
	"lej":        LET,
	"po":         TRUE,
	"jo":         FALSE,
	"nese":       IF,
	"perndryshe": ELSE,
	"kthe":       RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
