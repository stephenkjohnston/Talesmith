package parser

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/stephenkjohnston/talesmith/core/lexer"
	"github.com/stephenkjohnston/talesmith/core/scene"
)

type Parser struct {
	Tokens         []lexer.Token
	Position       int
	PlayerResponse string
	Scenes         []scene.Scene
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

func (p *Parser) Parse() error {
	for p.hasTokens() {
		token := p.advance()
		if token.Kind == lexer.FUNC_SCENE {
			if err := p.handleScene(); err != nil {
				return err
			}
		}

		if token.Kind == lexer.FUNC_ASK {
			p.handleAsk()
		}
	}

	return nil
}

func (p *Parser) handleScene() error {
	sceneToken := p.advance()

	if sceneToken.Kind != lexer.STRING {
		return errors.New("expected string after SCENE")
	}

	sceneName := p.extractQuotedContent(sceneToken.Literal)

	if err := p.sceneExists(sceneName); err != nil {
		return err
	}

	p.Scenes = append(p.Scenes, scene.Scene{Name: sceneName})
	return nil
}

func (p *Parser) sceneExists(name string) error {
	for _, scene := range p.Scenes {
		if scene.Name == name {
			return errors.New("duplicate scene name found \"" + name + "\"")
		}
	}

	return nil
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
