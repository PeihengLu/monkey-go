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

// nextToken returns the next token in the input string
func (l *Lexer) nextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
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
	case 0:
		// reached the end of the input
		tok.Literal = ""
		tok.Type = token.EOF
	}

	return tok
}

// newToken creates a new token with the given token type and literal
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}