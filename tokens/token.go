// tokens/token.go

package tokens

// creates a new type and attached underlying type
// string to it
// not as performant as int and byte, but easier to
// use since we can print it
type TokenType string

// go visibility is determined by the case of the 
// first letter of the identifier (function, struct
// and field)
// uppercase means exported (accessible from outside
// the package)
type Token struct {
	Type	TokenType
	Literal	string
}


// TokenTypes
const (
	ILLEGAL = "ILLEGAL" // unrecognizable tokens
	EOF		= "EOF"		// tells the parser that it can stop

	IDENT 	= "IDENT" // add, foobar, x, y etc.
	INT		= "INT"	  // 123456

	// operators
	ASSIGN	= "="
	PLUS 	= "+"

	// delimiters
	COMMA	  = ","
	SEMICOLON = ";"

	LPAREN	= "("
	RPAREN	= ")"
	LBRACKET = "{}"
	RBRACKET = "}"

	// keywords
	LET		 = "LET"
	FUNCTION = "FUNCTION"
)