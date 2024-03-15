package parse

type Token struct {
	Type  TokenType
	Value string
}

type TokenType int

const (
	tokenSection TokenType = iota
	tokenKey
	tokenValue
	tokenError
	tokenEOF
	tokenLeftBracket
	tokenRightBracket
	tokenEqual
	tokenNewline
)
