package lexer

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Tokenize(inputString string) []Token {
	var tokens []Token = []Token{}
	var ind int = 0
	for {
		var nextToken = parseNextToken(inputString, ind)
		if nextToken.Kind == SKIP {
			ind++
			continue
		}
		ind += nextToken.TokenLength
		tokens = append(tokens, nextToken)
		if nextToken.TokenLength == 0 {
			if nextToken.Kind != EOF {
				nextToken.Pos = ind // for showing error on position
			}
			break
		}
	}
	return tokens
}

func parseNextToken(inputString string, curInd int) Token {
	/*

		Hierarchy:
		1. EOF
		2. Whitespace
		3. keywords
		4. numbers (digit, -, .)
		5. string (quote)
		6. symbols
		7. error

	*/
	// EOF
	if curInd > len(inputString)-1 {
		return Token{
			Kind:        EOF,
			TokenLength: 0,
		}
	}
	var ch rune = rune(inputString[curInd])
	// whitespace
	if unicode.IsSpace(ch) {
		return Token{
			Kind: SKIP,
		}
	}
	// keywords
	extractedKeyword := extractKeyword(inputString, curInd)
	if extractedKeyword.Kind != UNKNOWN {
		return extractedKeyword
	}
	// numbers
	extractedNumber := extractNumber(inputString, curInd)
	if extractedNumber.Kind != UNKNOWN {
		return extractedNumber
	}
	// string
	extractedString := extractString(inputString, curInd)
	if extractedString.Kind != UNKNOWN {
		return extractedString
	}
	// symbols
	extractedSymbol := extractSymbol(inputString, curInd)
	if extractedSymbol.Kind != UNKNOWN {
		return extractedSymbol
	}
	// error
	return Token{
		Kind: UNKNOWN,
	}
}

func extractSymbol(inputString string, curInd int) Token {
	// symbols to check for
	// PIPE
	// SEMICOLON
	// COLON
	// UNDERSCORE
	// LBRACKET
	// RBRACKET
	// LPAREN
	// RPAREN
	// COMMA

	var ch rune = rune(inputString[curInd])
	switch ch {
	case PIPE.Rune():
		return Token{
			Kind:        PIPE,
			TokenLength: 1,
		}
	case SEMICOLON.Rune():
		return Token{
			Kind:        SEMICOLON,
			TokenLength: 1,
		}
	case COLON.Rune():
		return Token{
			Kind:        COLON,
			TokenLength: 1,
		}
	case UNDERSCORE.Rune():
		return Token{
			Kind:        UNDERSCORE,
			TokenLength: 1,
		}
	case LBRACKET.Rune():
		return Token{
			Kind:        LBRACKET,
			TokenLength: 1,
		}
	case RBRACKET.Rune():
		return Token{
			Kind:        RBRACKET,
			TokenLength: 1,
		}
	case LPAREN.Rune():
		return Token{
			Kind:        LPAREN,
			TokenLength: 1,
		}
	case RPAREN.Rune():
		return Token{
			Kind:        RPAREN,
			TokenLength: 1,
		}
	case COMMA.Rune():
		return Token{
			Kind:        COMMA,
			TokenLength: 1,
		}
	default:
		return Token{
			Kind: UNKNOWN,
		}
	}
}

func extractKeyword(inputString string, curInd int) Token {
	// tokens to check for
	// INT
	// BOOl
	// STRING
	// ARRAY
	// MAP
	// TRUE
	// FALSE

	var nextWord string = getNextWord(inputString, curInd)
	switch strings.ToLower(nextWord) {
	case INT.String():
		return Token{
			Kind:        INT,
			TokenLength: 3,
		}
	case BOOL.String():
		return Token{
			Kind:        BOOL,
			TokenLength: 4,
		}
	case STRING.String():
		return Token{
			Kind:        STRING,
			TokenLength: 6,
		}
	case ARRAY.String():
		return Token{
			Kind:        ARRAY,
			TokenLength: 5,
		}
	case MAP.String():
		return Token{
			Kind:        MAP,
			TokenLength: 3,
		}
	case TRUE.String():
		return Token{
			Kind:        TRUE,
			TokenLength: 4,
		}
	case FALSE.String():
		return Token{
			Kind:        FALSE,
			TokenLength: 5,
		}
	default:
		return Token{
			Kind: UNKNOWN,
		}
	}
}

func getNextWord(inputString string, curInd int) string {
	var curString strings.Builder
	for _, ch := range inputString[curInd:] {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			curString.WriteRune(ch)
		} else {
			break
		}
	}
	return curString.String()
}

func extractString(inputString string, curInd int) Token {
	if inputString[curInd] != '"' {
		return Token{
			Kind: UNKNOWN,
		}
		return Token{
			Kind: UNKNOWN,
		}
	}
	var stringToReturn strings.Builder
	var shouldEscape bool = false
	var didEnd bool = false
	for _, ch := range inputString[curInd+1:] {
		if ch == '"' {
			if !shouldEscape {
				didEnd = true
				break
			}
			shouldEscape = false
		} else if ch == '\\' {
			shouldEscape = true
			continue
		}
		stringToReturn.WriteRune(ch)
	}
	if didEnd {
		return Token{
			Kind:        STRING_LITERAL,
			Value:       stringToReturn.String(),
			TokenLength: stringToReturn.Len() + 2,
		}
	} else {
		return Token{
			Kind: UNKNOWN,
		}
	}
}

func extractNumber(inputString string, curInd int) Token {
	var stringNumber strings.Builder
	re := regexp.MustCompile("[-.0-9]")
	for _, ch := range inputString[curInd:] {
		if re.MatchString(string(ch)) {
			stringNumber.WriteRune(ch)
		} else {
			break
		}
	}
	if num, ok := strconv.ParseFloat(stringNumber.String(), 64); ok == nil {
		return Token{
			Kind:        NUMBER,
			TokenLength: stringNumber.Len(),
			Value:       num,
		}
	}
	return Token{
		Kind: UNKNOWN,
	}
}
