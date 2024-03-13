package parse

import "text/scanner"

type Token struct {
	Type  string
	Value string
}

type Lexer struct {
	Scanner scanner.Scanner
	Tokens  []string
	Pos     int
}
