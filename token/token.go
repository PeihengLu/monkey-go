// token/token.go

package token

// creates a new type and attached underlying type
// string to it
// not as performant as int and byte, but easier to
// use since we can print it
type TokenType string // the type definition only
// serves for better readability
// it is essentially just a string

// go visibility is determined by the case of the
// first letter of the identifier (function, struct
// and field)
// uppercase means exported (accessible from outside
// the package)
type Token struct {
	Type    TokenType
	Literal string
}

// TokenTypes
const (
	ILLEGAL = "ILLEGAL" // unrecognizable tokens
	EOF     = "EOF"     // tells the parser that it can stop

	IDENT = "IDENT" // add, foobar, x, y etc.
	INT   = "INT"   // 123456

	// operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	SLASH  = "/"
	BANG = "!" 
	ASTERISK = "*"	

	// logical operators
	LT = "<"
	GT = ">"

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"
	RETURN   = "RETURN"
	TRUE	 = "TRUE"
	FALSE	 = "FALSE"
	IF		 = "IF"
	ELSE	 = "ELSE"
)


// differentiates between keywords and identifiers using a map
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
	"return": RETURN,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
}

// LookupIdent checks if the identifier is a keyword
// and returns the keyword's TokenType
func LookupIdent(ident string) TokenType {
	// if the identifier is found in the keywords map
	// return the keyword's TokenType
	// the if statement is a shorthand for:
	// token, ok := keywords[ident]
	// if ok { ... }
	if token, ok := keywords[ident]; ok {
		return token
	}
	return IDENT
}