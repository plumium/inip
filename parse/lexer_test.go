package parse

import (
	"testing"
)

func TestSection(t *testing.T) {
	s := "[SectionName]"
	lex := NewLexer("", s)
	go lex.Run()
	for token := range lex.Token {
		t.Logf("type: %d, value: %s", token.Type, token.Value)
	}
}

func TestKeyValue(t *testing.T) {
	s := "Key=Value"
	lex := NewLexer("", s)
	go lex.Run()
	for token := range lex.Token {
		t.Logf("type: %d, value: %s", token.Type, token.Value)
	}
}

func TestValueWithBlank(t *testing.T) {
	s := "Key="
	lex := NewLexer("", s)
	go lex.Run()
	for token := range lex.Token {
		t.Logf("type: %d, value: %s", token.Type, token.Value)
	}
}

func TestIni(t *testing.T) {
	s := `
  [SectionName]
  key1=value1
  key2=value2
  `
	lex := NewLexer("", s)
	go lex.Run()
	for token := range lex.Token {
		t.Logf("type: %d, value: %s", token.Type, token.Value)
	}
}
