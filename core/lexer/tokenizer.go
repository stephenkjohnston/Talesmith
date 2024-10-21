package lexer

import (
	"fmt"
	"regexp"
	"strings"
)

type RegexHandler func(lex *Lexer, regex *regexp.Regexp)

type RegexPattern struct {
	Regex   *regexp.Regexp
	Handler RegexHandler
}

type Lexer struct {
	Tokens   []Token
	Source   string
	Position int
	Patterns []RegexPattern
}

/**
 * Whether we're at the end of the file
 * @return {bool}
 */
func (lex *Lexer) is_eof() bool {
	return lex.Position >= len(lex.Source)
}

func (lex *Lexer) advance(n int) {
	lex.Position += n
}

func (lex *Lexer) push(token Token) {
	lex.Tokens = append(lex.Tokens, token)
}

func (lex *Lexer) remainder() string {
	return lex.Source[lex.Position:]
}

/**
 * Creates a list of tokens
 * @return []Token
 */
func Tokenize(source string) []Token {
	lexer := NewLexer(source)

	for !lexer.is_eof() {
		matched := false

		for _, pattern := range lexer.Patterns {
			loc := pattern.Regex.FindStringIndex(lexer.remainder())

			if loc != nil && loc[0] == 0 {
				pattern.Handler(lexer, pattern.Regex)
				matched = true
				break
			}
		}

		if !matched {
			panic(fmt.Sprintf("Unrecognized token near %s\n", lexer.remainder()))
		}
	}

	lexer.push(NewToken(EOF, "EOF"))
	return lexer.Tokens
}

func defaultRegexHandler(kind TokenKind, literal string) RegexHandler {
	return func(lex *Lexer, regex *regexp.Regexp) {
		lex.advance(len(literal))
		lex.push(NewToken(kind, literal))
	}
}

func whitespaceHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	lex.advance(match[1])
}

func symbolHandler(lex *Lexer, regex *regexp.Regexp) {
	value := regex.FindString(lex.remainder())

	if kind, found := reservedKeywords[strings.ToLower(value)]; found {
		lex.push(NewToken(kind, value))
	} else {
		lex.push(NewToken(IDENTIFIER, value))
	}

	lex.advance(len(value))
}

func stringHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	stringLiteral := lex.remainder()[match[0]:match[1]]

	lex.push(NewToken(STRING, stringLiteral))
	lex.advance(len(stringLiteral))
}

/**
 * Creates a new instance of the lexer
 * @return Lexer
 */
func NewLexer(source string) *Lexer {
	return &Lexer{
		Position: 0,
		Source:   source,
		Tokens:   make([]Token, 0),
		Patterns: []RegexPattern{
			{regexp.MustCompile(`\s+`), whitespaceHandler},
			{regexp.MustCompile(`"[^"]*"`), stringHandler},
			{regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`), symbolHandler},
		},
	}
}
