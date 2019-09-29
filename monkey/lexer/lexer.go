package lexer

import "github.com/istsh/go-writing-an-interpreter/monkey/token"

type Lexer struct {
	input        string
	position     int  // 入力における現在の位置(現在の文字を指し示す)
	readPosition int  // これから読み込む位置(現在の文字の次)
	ch           byte // 現在検査中の文字
}

func New(input string) *Lexer {
	lexer := &Lexer{
		input: input,
	}

	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	// 次の文字が存在するか
	if lexer.readPosition >= len(lexer.input) {
		// 次の文字は存在しない(まだ何も読み込んでいない or ファイルの終わり)
		lexer.ch = 0
	} else {
		// 次の文字をセット
		lexer.ch = lexer.input[lexer.readPosition]
	}
	// 数値を1つ進める
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	switch lexer.ch {
	case '=':
		tok = newToken(token.ASSIGN, lexer.ch)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.ch)
	case '(':
		tok = newToken(token.LPAREN, lexer.ch)
	case ')':
		tok = newToken(token.RPAREN, lexer.ch)
	case ',':
		tok = newToken(token.COMMA, lexer.ch)
	case '+':
		tok = newToken(token.PLUS, lexer.ch)
	case '{':
		tok = newToken(token.LBRACE, lexer.ch)
	case '}':
		tok = newToken(token.RBRACE, lexer.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	lexer.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
