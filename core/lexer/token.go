package lexer

import (
	"fmt"
)

type TokenKind int

const (
	EOF TokenKind = iota
	IDENTIFIER
	SINGLE_QUOTE
	DOUBLE_QUOTE
	STRING

	// Reserved keywords
	FUNC_ASK
)

type Token struct {
	Kind    TokenKind
	Literal string
}

var reservedKeywords = map[string]TokenKind{
	"ask": FUNC_ASK,
}

func NewToken(kind TokenKind, literal string) Token {
	return Token{
		kind, literal,
	}
}

func (token Token) IsOneOf(expectedTokens ...TokenKind) bool {
	for _, expected := range expectedTokens {
		if expected == token.Kind {
			return true
		}
	}
	return false
}

func (token Token) Debug() {
	if token.IsOneOf(IDENTIFIER, STRING) {
		fmt.Printf("%s(%s)\n", token.TokenKindString(), token.Literal)
	} else {
		fmt.Printf("%s()\n", token.TokenKindString())
	}
}

func (token Token) TokenKindString() string {
	switch token.Kind {
	case EOF:
		return "eof"
	case STRING:
		return "string"
	case FUNC_ASK:
		return "func(ask)"
	default:
		return fmt.Sprintf("unknown(%d)", token.Kind)
	}
}
