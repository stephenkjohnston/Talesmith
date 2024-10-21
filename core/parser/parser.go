package parser

import (
	"fmt"
	"regexp"

	"github.com/stephenkjohnston/talesmith/core/lexer"
)

type Parser struct {
	Tokens         []lexer.Token
	Position       int
	PlayerResponse string
}

func (p *Parser) currentToken() lexer.Token {
	return p.Tokens[p.Position]
}

func (p *Parser) advance() lexer.Token {
	token := p.currentToken()
	p.Position++
	return token
}

func (p *Parser) hasTokens() bool {
	return p.Position < len(p.Tokens) && p.currentToken().Kind != lexer.EOF
}

func (p *Parser) nextToken() lexer.Token {
	return p.Tokens[p.Position+1]
}

func (p *Parser) previousToken() lexer.Token {
	return p.Tokens[p.Position-1]
}

func NewParser(source string) *Parser {
	tokens := lexer.Tokenize(source)

	return &Parser{
		Tokens:         tokens,
		Position:       0,
		PlayerResponse: "",
	}
}

func (p *Parser) Parse() {
	for p.hasTokens() {
		token := p.advance()
		if token.Kind == lexer.FUNC_ASK {
			p.handleAsk()
		}
	}
}

func (p *Parser) handleAsk() {
	questionToken := p.advance() // Move to the quoted content token

	if questionToken.Kind == lexer.STRING {
		content := p.extractQuotedContent(questionToken.Literal)
		fmt.Print(content + " ")
		var input string
		fmt.Scanln(&input)
		p.PlayerResponse = input
		fmt.Printf("You entered: %s\n", p.PlayerResponse)
	} else {
		fmt.Println("Syntax error: Expected a string after ASK.")
	}
}

func (p *Parser) extractQuotedContent(token string) string {
	re := regexp.MustCompile(`["'](.*?)["']`)
	match := re.FindStringSubmatch(token)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}
