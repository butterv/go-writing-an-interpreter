package parser

import (
	"github.com/istsh/go-writing-an-interpreter/ast"
	"github.com/istsh/go-writing-an-interpreter/lexer"
	"github.com/istsh/go-writing-an-interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// 2つのトークンを読み込む。curTokenとpeekTokenの両方がセットされる。
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
