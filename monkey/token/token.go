package token

type TokenType string

// 各文字列が何を意味しているのかを対応づける為に、定数を設けている。
const (
	ILLEGAL = "ILLEGAL" // 解析に失敗した場合に設定するTokenType
	EOF     = "EOF"     // コードの終了

	// 識別子 + リテラル
	IDENT  = "IDENT"  // 識別子 e.g. add, foobar, x, y, ...
	INT    = "INT"    // 数値 e.g. 1343456
	STRING = "STRING" // 文字列

	// 演算子
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	MACRO    = "MACRO"
)

type Token struct {
	Type    TokenType // トークンタイプ
	Literal string    // 実際の値
}

var keywaords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"macro":  MACRO,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywaords[ident]; ok {
		return tok
	}
	return IDENT
}
