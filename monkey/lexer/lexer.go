package lexer

import "github.com/istsh/go-writing-an-interpreter/monkey/token"

type Lexer struct {
	input        string // 入力
	position     int    // 入力における現在の位置(現在の文字を指し示す)
	readPosition int    // これから読み込む位置(現在の文字の次)
	ch           byte   // 現在検査中の文字
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}

	// 1文字進める
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// 空白、タブ、改行を飛ばす
	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	// 1文字進める
	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		// 空白、タブ、改行のときに飛ばして1文字進める
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	// 次の文字が存在するか
	if l.readPosition >= len(l.input) {
		// 次の文字は存在しない(まだ何も読み込んでいない or ファイルの終わり)
		l.ch = 0
	} else {
		// 次の文字をセット
		l.ch = l.input[l.readPosition]
	}
	// 数値を1つ進める
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	// 次の文字を覗き見る
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	// 識別子を抽出する
	position := l.position
	for isLetter(l.ch) {
		// 文字が途切れるまで読み込む
		l.readChar()
	}
	// positionから、readCharで進んだところまで抽出
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	// 数値を抽出する
	position := l.position
	for isDigit(l.ch) {
		// 文字が途切れるまで読み込む
		l.readChar()
	}
	// positionから、readCharで進んだところまで抽出
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	// 小文字大文字のアルファベットとアンダースコア
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	// 数値
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	// Tokenオブジェクトを初期化する
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (l *Lexer) readString() string {
	// 次のポジション
	position := l.position + 1
	for {
		// 1文字読み進める
		l.readChar()
		// ダブルクォートまたは0(終端)の場合はbreak
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}
