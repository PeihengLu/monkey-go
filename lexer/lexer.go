// lexer/lexer.go
package lexer

import "monkey-go/token"

// position and readPositions are initialized with default value
type Lexer struct {
	input        string
	position     int  // current position in input points to current character
	readPosition int  // current reading position in input after current character
	ch           byte // current char under examination
}

// New creates a new Lexer instance and returns a pointer to it
// dereferencing using * for a pointer is also possible, but go
// can handle the dereferencing when accessing an object's field
// such as l.position is the same as *l.position
func New(input string) *Lexer {
	// Create a new Lexer instance
	l := &Lexer{input: input}
	// Read the first character in the input string
	l.readChar()
	// Return a pointer to the Lexer instance
	return l
} 

// readChar reads the next character in the input string
func (l *Lexer) readChar() {
	// Check if the readPosition is greater than or equal to the length of the input string
	if l.readPosition >= len(l.input) {
		// Set the current character to 0 (NUL)
		l.ch = 0
	} else {
		// Set the current character to the character at the readPosition
		l.ch = l.input[l.readPosition]
	}
	// Set the position to the readPosition
	l.position = l.readPosition
	// Increment the readPosition
	l.readPosition++
}

// peekChar returns the next character in the input string
// without incrementing the readPosition
// used in the parsing of two-character operators
func (l *Lexer) peekChar() byte {
	// Check if the readPosition is greater than or equal to the length of the input string
	if l.readPosition >= len(l.input) {
		// Return 0 (NUL)
		return 0
	} else {
		// Return the character at the readPosition
		return l.input[l.readPosition]
	}
}

// nextToken returns the next token in the input string
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// skip whitespaces and newlines
	l.skipWhitespace()

	switch l.ch {
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		// check if the next character is an equal sign
		if l.peekChar() == '=' {
			// read the equal sign
			l.readChar()
			tok = newTokenString(token.NOT_EQ, "!=")
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '=':
		// check if the next character is an equal sign
		if l.peekChar() == '=' {
			// read the equal sign
			l.readChar()
			tok = newTokenString(token.EQ, "==")
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case 0:
		// reached the end of the input
		tok.Literal = ""	
		tok.Type = token.EOF
	default:
		// if the character is not a special character
		// we check if it is a letter
		if isLetter(l.ch) {
			// read the identifier
			tok.Literal = l.readIdentifier()
			// check if the identifier is a keyword
			tok.Type = token.LookupIdent(tok.Literal)
			// return the token to prevent the next readChar() call
			return tok
		} else if isDigit(l.ch) {
			// read the number
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			// return the token to prevent the next readChar() call
			return tok
		} else {
			// illegal token	
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// read the next character
	l.readChar()

	return tok
}

// skipWhitespace skips the whitespaces in the input string
func (l *Lexer) skipWhitespace() {
	// loop until the current character is not a whitespace
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		// read the next character
		l.readChar()
	}
}


// newToken creates a new token with the given token type and literal
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// overloaded newToken function that accepts a string literal
func newTokenString(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}

// readIdentifier reads the identifier from the input string
func (l *Lexer) readIdentifier() string {
	// reads the identifier until it encounters a non-letter character
	position := l.position
	// for loop in Go is similar to while loop in other languages
	// it continues until the condition is false
	// in this case, it continues until the character is not a letter
	// or an underscore
	for isLetter(l.ch) {
		l.readChar()
	}
	// return the identifier
	return l.input[position:l.position]
}

// readNumber reads the number from the input string
func (l *Lexer) readNumber() string {
	// reads the identifier until it encounters a non-digit char
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}
	// return the number
	return l.input[position: l.position]
}

// isLetter checks if the character is a letter
func isLetter(ch byte) bool {
	// check if the character is a letter or an underscore
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit checks if the character is a digit
func isDigit(ch byte) bool {
	// check if the character is a digit
	return '0' <= ch && ch <= '9'
}