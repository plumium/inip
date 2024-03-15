package parse

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type StateFn func(*Lexer) StateFn

type Lexer struct {
	input string
	state StateFn
	start int
	pos   int
	width int

	Name  string
	Token chan Token
}

const (
	eof          = -1
	leftBracket  = "["
	rightBracket = "]"
	equal        = "="
)

/*
Start a new lexer with a given input string. This returns the
instance of the lexer and a channel of tokens. Reading this stream
is the way to parse a given input and perform processing.
*/
func NewLexer(name string, input string) *Lexer {
	l := &Lexer{
		input: input,
		state: StateBegin,
		Name:  name,
		Token: make(chan Token, 3),
	}
	return l
}

/*
This lexer function starts everything off. It determines if we are
beginning with a key/value assignment or a section.
*/
func StateBegin(lex *Lexer) StateFn {
	lex.SkipWhitespace()
	if lex.IsEOF() {
		return nil
	}
	if strings.HasPrefix(lex.input[lex.pos:], leftBracket) {
		return StateLeftBracket
	} else {
		return StateKey
	}
}

/*
This lexer function emits a TOKEN_LEFT_BRACKET then returns
the lexer for a section header.
*/
func StateLeftBracket(lex *Lexer) StateFn {
	lex.pos += len(leftBracket)
	lex.Emit(tokenLeftBracket)
	return StateSection
}

/*
This lexer function emits a TOKEN_SECTION with the name of an
INI file section header.
*/
func StateSection(lex *Lexer) StateFn {
	for {
		if lex.IsEOF() {
			return lex.StateErrorf("Unexpected end of file")
		}
		if strings.HasPrefix(lex.input[lex.pos:], rightBracket) {
			lex.Emit(tokenSection)
			return StateRightBracket
		}
		lex.pos++
	}
}

/*
This lexer function emits a TOKEN_RIGHT_BRACKET then returns
the lexer for a begin.
*/
func StateRightBracket(lex *Lexer) StateFn {
	lex.pos += len(rightBracket)
	lex.Emit(tokenRightBracket)
	return StateBegin
}

/*
This lexer function emits a TOKEN_KEY with the name of an
key that will be assigned a value.
*/
func StateKey(lex *Lexer) StateFn {
	for {
		if strings.HasPrefix(lex.input[lex.pos:], equal) {
			lex.Emit(tokenKey)
			return StateEqualSign
		}
		lex.pos++
		if lex.IsEOF() {
			return lex.StateErrorf("Unexpected end of file")
		}
	}
}

/*
This lexer function emits a TOKEN_EQUAL_SIGN then returns
the lexer for value.
*/
func StateEqualSign(lex *Lexer) StateFn {
	lex.pos += len(equal)
	lex.Emit(tokenEqual)
	return StateValue
}

/*
This lexer function emits a TOKEN_VALUE with the value to be assigned to a key.
*/
func StateValue(lex *Lexer) StateFn {
	for {
		if strings.HasPrefix(lex.input[lex.pos:], "\n") {
			lex.Emit(tokenValue)
			return StateBegin
		}
		if lex.IsEOF() {
			lex.Emit(tokenValue)
			return nil
		}
		lex.pos++
	}
}

/*
Returns a token with error information.
*/
func (lex *Lexer) StateErrorf(format string, args ...interface{}) StateFn {
	lex.Token <- Token{
		Type:  tokenError,
		Value: fmt.Sprintf(format, args...),
	}
	lex.start = 0
	lex.pos = 0
	lex.width = 0
	return nil
}

/*
Starts the lexical analysis and feeding tokens into the token channel.
*/
func (lex *Lexer) Run() {
	for state := StateBegin; state != nil; {
		state = state(lex)
	}
	lex.Shutdown()
}

func (lex *Lexer) Next() rune {
	if lex.IsEOF() {
		return eof
	}
	r, w := utf8.DecodeRuneInString(lex.input[lex.pos:])
	lex.start = lex.pos
	lex.pos += w
	lex.width = w
	return r
}

func (lex *Lexer) IsEOF() bool {
	return lex.pos >= len(lex.input)
}

/*
Puts a token onto the token channel. The value of this token is
read from the input based on the current lexer position.
*/
func (lex *Lexer) Emit(tokenType TokenType) {
	lex.Token <- Token{
		Type:  tokenType,
		Value: lex.input[lex.start:lex.pos],
	}
	lex.start = lex.pos
}

/*
Backup to the beginning of the last read token.
*/
func (lex *Lexer) Backup() {
	lex.pos -= lex.width
}

/*
Skips whitespace until we get something meaningful.
*/
func (lex *Lexer) SkipWhitespace() {
	for {
		c := lex.Next()
		if c == eof {
			lex.Emit(tokenEOF)
			break
		}
		if !unicode.IsSpace(c) {
			lex.Backup()
			break
		}
	}
}

/*
Shuts down the token stream
*/
func (lex *Lexer) Shutdown() {
	close(lex.Token)
}
