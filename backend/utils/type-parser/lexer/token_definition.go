package lexer

type TokenKind int

type Token struct {
	Kind        TokenKind
	Value       any
	TokenLength int
	Pos         int
}

const (
	PIPE TokenKind = iota
	SEMICOLON
	COLON
	UNDERSCORE
	LBRACKET
	RBRACKET
	LPAREN
	RPAREN
	COMMA

	INT
	BOOL
	STRING
	ARRAY
	MAP
	TRUE
	FALSE

	NUMBER
	STRING_LITERAL

	EOF

	UNKNOWN
	SKIP
	/*
		main tokens:
		1. PIPE        |
		2. SEMICOLON   ;
		3. COLON       :
		4. UNDERSCORE  _
		5. LBRACKET    [
		6. RBRACKET    ]
		7. LPAREN      (
		8. RPAREN      )
		9. COMMA       ,

		identity tokens:
		1. INT
		2. BOOL
		3. STRING
		4. ARRAY
		5. MAP
		6. TRUE
		7. FALSE

		tokens with value:
		1. NUMBER(value)
		2. STRING_LITERAL(value)

		indicator tokens:
		1. EOF

		error tokens:
		1. UNKNOWN
		2. SKIP
	*/
)

var tokenName = map[TokenKind]string{
	PIPE:           "|",
	SEMICOLON:      ";",
	COLON:          ":",
	UNDERSCORE:     "_",
	LBRACKET:       "[",
	RBRACKET:       "]",
	LPAREN:         "(",
	RPAREN:         ")",
	COMMA:          ",",
	NUMBER:         "NUMBER",
	STRING_LITERAL: "\"",
	INT:            "int",
	BOOL:           "bool",
	STRING:         "string",
	ARRAY:          "array",
	MAP:            "map",
	TRUE:           "true",
	FALSE:          "false",
	EOF:            "EOF",
	UNKNOWN:        "UNKNOWN",
	SKIP:           "SKIP",
}

func (token TokenKind) String() string {
	if name, ok := tokenName[token]; ok {
		return name
	}

	return tokenName[UNKNOWN]
}

func (token TokenKind) Rune() rune {
	if name, ok := tokenName[token]; ok {
		return rune(name[0])
	}

	return rune(tokenName[UNKNOWN][0])
}
