package language

import (
	"fmt"

	"strconv"

	"github.com/ijsnow/goql/internal/errors"
	"github.com/ijsnow/goql/internal/language"
)

// CreateLexer returns a Lexer for that source given a Source object
// A Lexer is a stateful stream generator in that every time
// it is advanced, it returns the next token in the Source. Assuming the
// source lexes, the final Token emitted by the lexer will be of kind
// EOF, after which the lexer will repeatedly return the same EOF token
// whenever called.
func CreateLexer(source language.Source) Lexer {
	startOfFileToken := *language.NewToken(language.TokenSOF, 0, 0, 0, 0, nil, "")

	return Lexer{
		Source:    source,
		LastToken: startOfFileToken,
		Token:     startOfFileToken,
		Line:      1,
		LineStart: 0,
	}
}

// Advance advances the lexer to the next token we are interested in
func (l Lexer) Advance() (*language.Token, error) {
	token := l.Token
	l.LastToken = l.Token

	for token.Kind != language.TokenEOF {
		t, err := readToken(l, &token)
		if err != nil {
			return nil, err
		}
		token = *t
		token.Next = t

		if token.Kind != language.TokenComment {
			break
		}
	}

	l.Token = token

	return &token, nil
}

// Lexer is the return type of createLexer
type Lexer struct {
	Source language.Source

	/**
	 * The previously focused non-ignored token.
	 */
	LastToken language.Token

	/**
	 * The currently focused non-ignored token.
	 */
	Token language.Token

	/**
	 * The (1-indexed) line containing the current token.
	 */
	Line int

	/**
	 * The character offset at which the current line begins.
	 */
	LineStart int
}

// GetTokenDesc is a helper function to describe a token as a string for debugging
func GetTokenDesc(token language.Token) string {
	value := token.Value

	if value == "" {
		return string(token.Kind)
	}

	return fmt.Sprintf(`%s "%s"`, string(token.Kind), value)
}

// ReadToken gets the next token from the source starting at the given position.
//
// This skips over whitespace and comments until it finds the next lexable
// token, then lexes punctuators immediately or calls the appropriate helper
// function for more complicated tokens.
func readToken(lexer Lexer, prev *language.Token) (*language.Token, error) {
	source := lexer.Source
	body := source.Body
	bodyLength := len(body)

	position := positionAfterWhitespace(body, prev.End, &lexer)
	line := lexer.Line
	col := 1 + position - lexer.LineStart

	if position >= bodyLength {
		return language.NewToken(language.TokenEOF, bodyLength, bodyLength, line, col, prev, ""), nil
	}

	code := charCodeAt(body, position)

	// SourceCharacter
	if code < 0x0020 && code != 0x0009 && code != 0x000A && code != 0x000D {
		return nil, errors.NewSyntaxError(
			source,
			position,
			fmt.Sprintf("Cannot contain the invalid character %s.", printCharCode(code)),
		)
	}

	switch code {
	// !
	case 33:
		return language.NewToken(language.TokenBang, position, position+1, line, col, prev, ""), nil
	// #
	case 35:
		return readComment(source, position, line, col, prev), nil
	// $
	case 36:
		return language.NewToken(language.TokenDollar, position, position+1, line, col, prev, ""), nil
	// (
	case 40:
		return language.NewToken(language.TokenParenLeft, position, position+1, line, col, prev, ""), nil
	// )
	case 41:
		return language.NewToken(language.TokenParenRight, position, position+1, line, col, prev, ""), nil
	// . -> (...)
	case 46:
		if charCodeAt(body, position+1) == 46 && charCodeAt(body, position+2) == 46 {
			return language.NewToken(language.TokenSpread, position, position+3, line, col, prev, ""), nil
		}

	// :
	case 58:
		return language.NewToken(language.TokenColon, position, position+1, line, col, prev, ""), nil
	// =
	case 61:
		return language.NewToken(language.TokenEqual, position, position+1, line, col, prev, ""), nil
	// @
	case 64:
		return language.NewToken(language.TokenAt, position, position+1, line, col, prev, ""), nil
	// [
	case 91:
		return language.NewToken(language.TokenBracketLeft, position, position+1, line, col, prev, ""), nil
	// ]
	case 93:
		return language.NewToken(language.TokenBracketRight, position, position+1, line, col, prev, ""), nil
	// {
	case 123:
		return language.NewToken(language.TokenBraceLeft, position, position+1, line, col, prev, ""), nil
	// |
	case 124:
		return language.NewToken(language.TokenPipe, position, position+1, line, col, prev, ""), nil
	// }
	case 125:
		return language.NewToken(language.TokenBraceRight, position, position+1, line, col, prev, ""), nil
	// A-Z _ a-z
	case 65,
		66,
		67,
		68,
		69,
		70,
		71,
		72,
		73,
		74,
		75,
		76,
		77,
		78,
		79,
		80,
		81,
		82,
		83,
		84,
		85,
		86,
		87,
		88,
		89,
		90,
		95,
		97,
		98,
		99,
		100,
		101,
		102,
		103,
		104,
		105,
		106,
		107,
		108,
		109,
		110,
		111,
		112,
		113,
		114,
		115,
		116,
		117,
		118,
		119,
		120,
		121,
		122:
		return readName(source, position, line, col, prev), nil
	// - 0-9
	case 45,
		48,
		49,
		50,
		51,
		52,
		53,
		54,
		55,
		56,
		57:
		return readNumber(source, position, code, line, col, prev)
	// "
	case 34:
		return readString(source, position, line, col, prev)
	}

	return nil, errors.NewSyntaxError(
		source,
		position,
		unexpectedCharacterMessage(code),
	)
}

func charCodeAt(s string, n int) rune {
	i := 0
	for _, r := range s {
		if i == n {
			return r
		}
		i++
	}
	return -1
}

func sliceStr(str string, start int, pos int) string {
	return string([]rune(str)[start:pos])
}

func printCharCode(code rune) string {
	// NaN/undefined represents access beyond the end of the file.
	if code == -1 {
		return string(language.TokenEOF)
	}

	ucode := fmt.Sprintf(`%#U`, code)

	// convert unicode format from U+XXXX to \uXXXX
	return fmt.Sprintf(`\u%s"`, ucode[2:len(ucode)])
}

func printChar(code rune) string {
	// NaN/undefined represents access beyond the end of the file.
	if code == -1 {
		return string(language.TokenEOF)
	}

	// EX: is ? not \u003F
	if strconv.IsGraphic(code) {
		return fmt.Sprintf("\"%c\"", code)
	}

	return printCharCode(code)
}

// unexpectedCharacterMessage reports a message that an unexpected character was encountered.
func unexpectedCharacterMessage(code rune) string {
	if code == 39 { // '
		return "Unexpected single quote character ('), did you mean to use a double quote (\")?"
	}

	q := fmt.Sprintf("%+q", code)

	//return fmt.Sprintf("Cannot parse the unexpected character %s.", printChar(code))
	return fmt.Sprintf("Cannot parse the unexpected character \"%s\".", q[1:len(q)-1])
}

/**
 * positionAfterWhitespace reads from body starting at startPosition until it finds a non-whitespace
 * or commented character, then returns the position of that character for
 * lexing.
 */
func positionAfterWhitespace(body string, startPosition int, lexer *Lexer) int {
	bodyLength := len(body)
	position := startPosition

	for position < bodyLength {
		code := charCodeAt(body, position)
		// tab | space | comma | BOM
		if code == 9 || code == 32 || code == 44 || code == 0xFEFF {
			position++
		} else if code == 10 { // new line
			position++
			lexer.Line++
			lexer.LineStart = position
		} else if code == 13 { // carriage return
			if charCodeAt(body, position+1) == 10 {
				position += 2
			} else {
				position++
			}
			lexer.Line++
			lexer.LineStart = position
		} else {
			break
		}
	}

	return position
}

/* readComment reads a comment token from the source file.
 *
 * #[\u0009\u0020-\uFFFF]*
 */
func readComment(source language.Source, start, line, col int, prev *language.Token) *language.Token {
	var code rune
	body := source.Body
	position := start

	position++
	code = charCodeAt(body, position)
	for code != 0 &&
		// SourceCharacter but not LineTerminator
		(code > 0x001F || code == 0x0009) {
		position++
		code = charCodeAt(body, position)
	}

	return language.NewToken(
		language.TokenComment,
		start,
		position,
		line,
		col,
		prev,
		sliceStr(body, start+1, position),
	)
}

/**
 * Reads a number token from the source file, either a float
 * or an int depending on whether a decimal point appears.
 *
 * Int:   -?(0|[1-9][0-9]*)
 * Float: -?(0|[1-9][0-9]*)(\.[0-9]+)?((E|e)(+|-)?[0-9]+)?
 */
func readNumber(
	source language.Source,
	start int,
	firstCode rune,
	line,
	col int,
	prev *language.Token,
) (*language.Token, error) {
	body := source.Body
	code := firstCode
	position := start
	isFloat := false

	var err error

	if code == 45 { // -
		position++
		code = charCodeAt(body, position)
	}

	if code == 48 { // 0
		position++
		code = charCodeAt(body, position)
		if code >= 48 && code <= 57 {
			return nil, errors.NewSyntaxError(
				source,
				position,
				fmt.Sprintf("Invalid number, unexpected digit after 0: \"%c\".", code),
			)
		}
	} else {
		position, err = readDigits(source, position, code)
		if err != nil {
			return nil, err
		}

		code = charCodeAt(body, position)
	}

	if code == 46 { // .
		isFloat = true
		position++
		code = charCodeAt(body, position)
		position, err = readDigits(source, position, code)
		if err != nil {
			return nil, err
		}
		code = charCodeAt(body, position)
	}

	if code == 69 || code == 101 { // E e
		isFloat = true
		position++
		code = charCodeAt(body, position)
		if code == 43 || code == 45 { // + -
			position++
			code = charCodeAt(body, position)
		}
		position, err = readDigits(source, position, code)
		if err != nil {
			return nil, err
		}
	}

	var typeToken language.TokenKind
	if isFloat {
		typeToken = language.TokenFloat
	} else {
		typeToken = language.TokenInt
	}

	return language.NewToken(
		typeToken,
		start,
		position,
		line,
		col,
		prev,
		sliceStr(body, start, position),
	), nil
}

/* readDigits returns the new position in the source after reading digits.
 */
func readDigits(source language.Source, start int, firstCode rune) (int, error) {
	body := source.Body
	position := start
	code := firstCode

	if code >= 48 && code <= 57 { // 0 - 9
		position++
		code = charCodeAt(body, position)
		for code >= 48 && code <= 57 { // 0 - 9
			position++
			code = charCodeAt(body, position)
		}
		return position, nil
	}

	return position, errors.NewSyntaxError(
		source,
		position,
		fmt.Sprintf("Invalid number, expected digit but got: %s.", printChar(code)),
	)
}

/**
 * Reads a string token from the source file.
 *
 * "([^"\\\u000A\u000D]|(\\(u[0-9a-fA-F]{4}|["\\/bfnrt])))*"
 */
func readString(source language.Source, start, line, col int, prev *language.Token) (*language.Token, error) {
	body := source.Body
	position := start + 1
	chunkStart := position
	var code rune
	value := ""

	for position < len(body) {
		code = charCodeAt(body, position)

		if code != -1 &&
			// not LineTerminator
			code != 0x000A && code != 0x000D &&
			// not Quote (")
			code != 34 {
			// SourceCharacter
			if code < 0x0020 && code != 0x0009 {
				return nil, errors.NewSyntaxError(
					source,
					position,
					fmt.Sprintf("Invalid character within String: %s.", printCharCode(code)),
				)
			}

			position++
			if code == 92 { // \
				value += sliceStr(body, chunkStart, position-1)
				code = charCodeAt(body, position)
				switch code {
				case 34:
					value += "\""
				case 47:
					value += "/"
				case 92:
					value += "\\"
				case 98:
					value += "\b"
				case 102:
					value += "\f"
				case 110:
					value += "\n"
				case 114:
					value += "\r"
				case 116:
					value += "\t"
				case 117: // u
					charCode := uniCharCode(
						charCodeAt(body, position+1),
						charCodeAt(body, position+2),
						charCodeAt(body, position+3),
						charCodeAt(body, position+4),
					)

					if charCode < 0 {
						return nil, errors.NewSyntaxError(
							source,
							position,
							fmt.Sprintf("Invalid character escape sequence: \\u%s.",
								sliceStr(body, position+1, position+5)),
						)
					}
					value += fmt.Sprintf("%c", charCode)
					position += 4
				default:
					return nil, errors.NewSyntaxError(
						source,
						position,
						fmt.Sprintf("Invalid character escape sequence: \\%c.", code),
					)
				}

				position++
				chunkStart = position
			}
		} else {
			break
		}
	}

	if code != 34 { // quote (")
		return nil, errors.NewSyntaxError(source, position, "Unterminated string.")
	}

	value += sliceStr(body, chunkStart, position)
	return language.NewToken(language.TokenString, start, position+1, line, col, prev, value), nil
}

/**
 * Converts four hexidecimal chars to the integer that the
 * string represents. For example, uniCharCode('0','0','0','f')
 * will return 15, and uniCharCode('0','0','f','f') returns 255.
 *
 * Returns a negative number on error, if a char was invalid.
 *
 * This is implemented by noting that char2hex() returns -1 on error,
 * which means the result of ORing the char2hex() will also be negative.
 */
func uniCharCode(a, b, c, d rune) rune {
	return rune(char2hex(a)<<12 | char2hex(b)<<8 | char2hex(c)<<4 | char2hex(d))
}

/**
 * Converts a hex character to its integer value.
 * '0' becomes 0, '9' becomes 9
 * 'A' becomes 10, 'F' becomes 15
 * 'a' becomes 10, 'f' becomes 15
 *
 * Returns -1 on error.
 */
func char2hex(a rune) int {
	if a >= 48 && a <= 57 { // 0-9
		return int(a - 48)
	} else if a >= 65 && a <= 70 { // A-F
		return int(a - 55)
	} else if a >= 97 && a <= 102 { // a-f
		return int(a - 87)
	}

	return -1
}

/**
 * Reads an alphanumeric + underscore name from the source.
 *
 * [_A-Za-z][_0-9A-Za-z]*
 */
func readName(source language.Source, position, line, col int, prev *language.Token) *language.Token {
	body := source.Body
	bodyLength := len(body)
	end := position + 1
	var code rune

	if end != bodyLength {
		code = charCodeAt(body, end)

		for end != bodyLength && code != -1 && (code == 95 || // _
			(code >= 48 && code <= 57) || // 0-9
			(code >= 65 && code <= 90) || // A-Z
			(code >= 97 && code <= 122)) {
			end++
			code = charCodeAt(body, end)
		}
	}

	return language.NewToken(
		language.TokenName,
		position,
		end,
		line,
		col,
		prev,
		sliceStr(body, position, end),
	)
}
