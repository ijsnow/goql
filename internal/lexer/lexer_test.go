package language

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/ijsnow/goql/internal/language"
)

func lexOne(str string) (*language.Token, error) {
	lexer := CreateLexer(language.NewSource(str))

	return lexer.Advance()
}

func TestDisallowsUncommonControlCharacters(t *testing.T) {
	_, err := lexOne("\u0007")
	if err == nil {
		t.Errorf("lexer allowed uncommon control chars %v", err)
	}

	want := `Syntax Error GraphQL request (1:1) Cannot contain the invalid character "\u0007".

1: 
   ^
`

	if err.Error() != want {
		t.Errorf("error was formatted incorrectly\ngot: \n%vwanted: \n%v", err.Error(), want)
	}
}

func TestAcceptsBOMHeader(t *testing.T) {
	got, err := lexOne("\uFEFF foo")
	if err != nil {
		t.Fatal(err)
	}

	if got.Kind != language.TokenName {
		t.Errorf("unexpected token kind; got %v wanted %v", got.Kind, language.TokenName)
	}

	if got.Start != 2 {
		t.Errorf("unexpected token start; got %v wanted %v", got.Start, 2)
	}

	if got.End != 5 {
		t.Errorf("unexpected token end; got %v wanted %v", got.End, 5)
	}

	if got.Value != "foo" {
		t.Errorf("unexpected token value; got %v wanted %v", got.Value, "foo")
	}
}
func TestRecordsLineAndColumn(t *testing.T) {
	got, err := lexOne("\n \r\n \r  foo\n")
	if err != nil {
		t.Fatal(err)
	}

	if got.Kind != language.TokenName {
		t.Errorf("unexpected token kind; got %v wanted %v", got.Kind, language.TokenName)
	}

	if got.Start != 8 {
		t.Errorf("unexpected token start; got %v wanted %v", got.Start, 8)
	}

	if got.End != 11 {
		t.Errorf("unexpected token end; got %v wanted %v", got.End, 11)
	}

	if got.Line != 4 {
		t.Errorf("unexpected token line; got %v wanted %v", got.Line, 4)
	}

	if got.Column != 3 {
		t.Errorf("unexpected token column; got %v wanted %v", got.End, 3)
	}

	if got.Value != "foo" {
		t.Errorf("unexpected token value; got %v wanted %v", got.Value, "foo")
	}
}

func TestCanSerialize(t *testing.T) {
	token, err := lexOne("foo")
	if err != nil {
		t.Fatal(err)
	}

	b, err := json.Marshal(token)
	if err != nil {
		t.Fatal(err)
	}

	got := string(b)

	// Ordered by struct definition.
	// TODO: Do they always stay in the order they are defined???
	want := `{"kind":"Name","line":1,"column":1,"value":"foo"}`

	if got != want {
		t.Errorf("json serialization yielded unexpected result; got\n%v\nwanted\n%v", got, want)
	}
}

func TestSkipsWhitespaceAndComments(t *testing.T) {
	got, err := lexOne(`

    foo

`)
	if err != nil {
		t.Fatal(err)
	}

	if got.Kind != language.TokenName {
		t.Errorf("kind: got %v wanted %v", got.Kind, language.TokenName)
	}

	if got.Start != 6 {
		t.Errorf("kind: got %v wanted %v", got.Start, 6)
	}

	if got.End != 9 {
		t.Errorf("kind: got %v wanted %v", got.End, 9)
	}

	if got.Value != "foo" {
		t.Errorf("kind: got %v wanted %v", got.Value, "foo")
	}

	got, err = lexOne(`
    #comment
    foo#comment
`)
	if err != nil {
		t.Fatal(err)
	}

	if got.Kind != language.TokenName {
		t.Errorf("kind: got %v wanted %v", got.Kind, language.TokenName)
	}

	if got.Start != 18 {
		t.Errorf("kind: got %v wanted %v", got.Start, 18)
	}

	if got.End != 21 {
		t.Errorf("kind: got %v wanted %v", got.End, 21)
	}

	if got.Value != "foo" {
		t.Errorf("kind: got %v wanted %v", got.Value, "foo")
	}

	got, err = lexOne(",,,foo,,,")
	if err != nil {
		t.Fatal(err)
	}

	if got.Kind != language.TokenName {
		t.Errorf("kind: got %v wanted %v", got.Kind, language.TokenName)
	}

	if got.Start != 3 {
		t.Errorf("start: got %v wanted %v", got.Start, 3)
	}

	if got.End != 6 {
		t.Errorf("end: got %v wanted %v", got.End, 6)
	}

	if got.Value != "foo" {
		t.Errorf("value: got %v wanted %v", got.Value, "foo")
	}
}

func TestErrorsRespectWhitespace(t *testing.T) {
	test := `

    ?

`
	_, err := lexOne(test)
	if err == nil {
		t.Errorf("did not get an error for %v", test)
	}

	want := "Syntax Error GraphQL request (3:5) " +
		"Cannot parse the unexpected character \"?\".\n" +
		"\n" +
		"2: \n" +
		"3:     ?\n" +
		"       ^\n" +
		"4: \n"

	if err.Error() != want {
		t.Errorf("got unexpected error;\ngot\n%v\nwanted\n%v", err.Error(), want)
	}
}

type tokenTest struct {
	lex  string
	want *language.Token
}

func checkToken(t *testing.T, input string, want *language.Token) {
	got, err := lexOne(input)
	if err != nil {
		t.Fatalf("%v err %v", input, err)
	}

	checkExistingToken(t, input, got, want)
}

func checkExistingToken(t *testing.T, input string, got *language.Token, want *language.Token) {
	if got.Kind != want.Kind {
		t.Errorf("kind %v: got %v wanted %v", input, got.Kind, want.Kind)
	}

	if got.Start != want.Start {
		t.Errorf("start %v: got %v wanted %v", input, got.Start, want.Start)
	}

	if got.End != want.End {
		t.Errorf("end %v: got %v wanted %v", input, got.End, want.End)
	}

	if got.Value != want.Value {
		t.Errorf("value %v: got %v wanted %v", input, got.Value, want.Value)
	}
}

func TestLexesStrings(t *testing.T) {
	set := []tokenTest{
		tokenTest{
			lex: "\"simple\"",
			want: &language.Token{
				Kind:  language.TokenString,
				Start: 0,
				End:   8,
				Value: "simple",
			},
		},
		tokenTest{
			lex: `" white space "`,
			want: &language.Token{
				Kind:  language.TokenString,
				Start: 0,
				End:   15,
				Value: " white space ",
			},
		},
		tokenTest{
			lex: "\"quote \\\"\"",
			want: &language.Token{
				Kind:  language.TokenString,
				Start: 0,
				End:   10,
				Value: `quote "`,
			},
		},
		tokenTest{
			lex: "\"escaped \\n\\r\\b\\t\\f\"",
			want: &language.Token{
				Kind:  language.TokenString,
				Start: 0,
				End:   20,
				Value: "escaped \n\r\b\t\f",
			},
		},
		tokenTest{
			lex: "\"slashes \\\\ \\/\"",
			want: &language.Token{
				Kind:  language.TokenString,
				Start: 0,
				End:   15,
				Value: "slashes \\ /",
			},
		},
		tokenTest{
			lex: "\"unicode \\u1234\\u5678\\u90AB\\uCDEF\"",
			want: &language.Token{
				Kind:  language.TokenString,
				Start: 0,
				End:   34,
				Value: "unicode \u1234\u5678\u90AB\uCDEF",
			},
		},
	}

	for _, test := range set {
		checkToken(t, test.lex, test.want)
	}
}

func testErr(t *testing.T, err error, want string) {
	if err == nil {
		t.Errorf("expected error but got none (wanted %v)", want)
		return
	}

	if !strings.Contains(strings.ToLower(err.Error()), strings.ToLower(want)) {
		t.Errorf("wanted err to contain\n%v, got\n%v", want, err.Error())
	}
}

func TestLexReportsUsefulStringErrors(t *testing.T) {
	set := [][]string{
		[]string{
			"\"",
			"Syntax Error GraphQL request (1:2) Unterminated string.",
		},
		[]string{
			"\"",
			"Syntax Error GraphQL request (1:2) Unterminated string.",
		},
		[]string{
			"\"no end quote",
			"Syntax Error GraphQL request (1:14) Unterminated string.",
		},
		[]string{
			"'single quotes'",
			"Syntax Error GraphQL request (1:1) Unexpected single quote character ('), " +
				"did you mean to use a double quote (\")?",
		},
		[]string{
			"\"contains unescaped \u0007 control char\"",
			"Syntax Error GraphQL request (1:21) Invalid character within String: \"\\u0007\".",
		},
		[]string{
			"\"null-byte is not \u0000 end of file\"",
			"Syntax Error GraphQL request (1:19) Invalid character within String: \"\\u0000\".",
		},
		[]string{
			"\"multi\nline\"",
			"Syntax Error GraphQL request (1:7) Unterminated string",
		},
		[]string{
			"\"multi\rline\"",
			"Syntax Error GraphQL request (1:7) Unterminated string",
		},
		[]string{
			"\"bad \\z esc\"",
			"Syntax Error GraphQL request (1:7) Invalid character escape sequence: \\z.",
		},
		[]string{
			"\"bad \\x esc\"",
			"Syntax Error GraphQL request (1:7) Invalid character escape sequence: \\x.",
		},
		[]string{
			"\"bad \\u1 esc\"",
			"Syntax Error GraphQL request (1:7) Invalid character escape sequence: \\u1 es.",
		},
		[]string{
			"\"bad \\u0XX1 esc\"",
			"Syntax Error GraphQL request (1:7) Invalid character escape sequence: \\u0XX1.",
		},
		[]string{
			"\"bad \\uXXXX esc\"",
			"Syntax Error GraphQL request (1:7) Invalid character escape sequence: \\uXXXX.",
		},
		[]string{
			"\"bad \\uFXXX esc\"",
			"Syntax Error GraphQL request (1:7) Invalid character escape sequence: \\uFXXX.",
		},
		[]string{
			"\"bad \\uXXXF esc\"",
			"Syntax Error GraphQL request (1:7) Invalid character escape sequence: \\uXXXF.",
		},
	}

	for _, test := range set {
		_, err := lexOne(test[0])
		testErr(t, err, test[1])
	}

}

func TestLexesNumbers(t *testing.T) {
	set := []tokenTest{
		tokenTest{
			lex: "4",
			want: &language.Token{
				Kind:  language.TokenInt,
				Start: 0,
				End:   1,
				Value: "4",
			},
		},
		tokenTest{
			lex: "4.123",
			want: &language.Token{
				Kind:  language.TokenFloat,
				Start: 0,
				End:   5,
				Value: "4.123",
			},
		},
		tokenTest{
			lex: "-4",
			want: &language.Token{
				Kind:  language.TokenInt,
				Start: 0,
				End:   2,
				Value: "-4",
			},
		},
		tokenTest{
			lex: "9",
			want: &language.Token{
				Kind:  language.TokenInt,
				Start: 0,
				End:   1,
				Value: "9",
			},
		},
		tokenTest{
			lex: "0",
			want: &language.Token{
				Kind:  language.TokenInt,
				Start: 0,
				End:   1,
				Value: "0",
			},
		},
		tokenTest{
			lex: "-4.123",
			want: &language.Token{
				Kind:  language.TokenFloat,
				Start: 0,
				End:   6,
				Value: "-4.123",
			},
		},
		tokenTest{
			lex: "0.123",
			want: &language.Token{
				Kind:  language.TokenFloat,
				Start: 0,
				End:   5,
				Value: "0.123",
			},
		},

		tokenTest{
			lex: "123e4",
			want: &language.Token{
				Kind:  language.TokenFloat,
				Start: 0,
				End:   5,
				Value: "123e4",
			},
		},

		tokenTest{
			lex: "123E4",
			want: &language.Token{
				Kind:  language.TokenFloat,
				Start: 0,
				End:   5,
				Value: "123E4",
			},
		},

		tokenTest{
			lex: "123e-4",
			want: &language.Token{
				Kind:  language.TokenFloat,
				Start: 0,
				End:   6,
				Value: "123e-4",
			},
		},
		tokenTest{
			lex: "123e+4",
			want: &language.Token{
				Kind:  language.TokenFloat,
				Start: 0,
				End:   6,
				Value: "123e+4",
			},
		},
		tokenTest{
			lex: "-1.123e4",
			want: &language.Token{
				Kind:  language.TokenFloat,
				Start: 0,
				End:   8,
				Value: "-1.123e4",
			},
		},
		tokenTest{
			lex: "-1.123E4",
			want: &language.Token{
				Kind:  language.TokenFloat,
				Start: 0,
				End:   8,
				Value: "-1.123E4",
			},
		},

		tokenTest{
			lex: "-1.123e-4",
			want: &language.Token{
				Kind:  language.TokenFloat,
				Start: 0,
				End:   9,
				Value: "-1.123e-4",
			},
		},
		tokenTest{
			lex: "-1.123e+4",
			want: &language.Token{
				Kind:  language.TokenFloat,
				Start: 0,
				End:   9,
				Value: "-1.123e+4",
			},
		},
		tokenTest{
			lex: "-1.123e4567",
			want: &language.Token{
				Kind:  language.TokenFloat,
				Start: 0,
				End:   11,
				Value: "-1.123e4567",
			},
		},
	}

	for _, test := range set {
		checkToken(t, test.lex, test.want)
	}
}

func TestLexReportsUserfulNumberErrors(t *testing.T) {
	set := [][]string{
		[]string{
			"00",
			"Syntax Error GraphQL request (1:2) Invalid number, " +
				"unexpected digit after 0: \"0\".",
		},
		[]string{
			"00",
			"Syntax Error GraphQL request (1:2) Invalid number, " +
				"unexpected digit after 0: \"0\".",
		},
		[]string{
			"+1",
			"Syntax Error GraphQL request (1:1) Cannot parse the unexpected character \"+\".",
		},
		[]string{
			"1.",
			"Syntax Error GraphQL request (1:3) Invalid number, " +
				"expected digit but got: <EOF>.",
		},
		[]string{
			".123",
			"Syntax Error GraphQL request (1:1) Cannot parse the unexpected character \".\".",
		},
		[]string{
			"1.A",
			"Syntax Error GraphQL request (1:3) Invalid number, " +
				"expected digit but got: \"A\".",
		},
		[]string{
			"-A",
			"Syntax Error GraphQL request (1:2) Invalid number, " +
				"expected digit but got: \"A\".",
		},
		[]string{
			"1.0e",
			"Syntax Error GraphQL request (1:5) Invalid number, " +
				"expected digit but got: <EOF>.",
		},
		[]string{
			"1.0eA",
			"Syntax Error GraphQL request (1:5) Invalid number, " +
				"expected digit but got: \"A\".",
		},
	}

	for _, test := range set {
		_, err := lexOne(test[0])
		testErr(t, err, test[1])
	}
}

func TestLexesPuncuation(t *testing.T) {
	set := []tokenTest{
		tokenTest{
			lex: "!",
			want: &language.Token{
				Kind:  language.TokenBang,
				Start: 0,
				End:   1,
				Value: "",
			},
		},
		tokenTest{
			lex: "$",
			want: &language.Token{
				Kind:  language.TokenDollar,
				Start: 0,
				End:   1,
				Value: "",
			},
		},
		tokenTest{
			lex: "(",
			want: &language.Token{
				Kind:  language.TokenParenLeft,
				Start: 0,
				End:   1,
				Value: "",
			},
		},
		tokenTest{
			lex: ")",
			want: &language.Token{
				Kind:  language.TokenParenRight,
				Start: 0,
				End:   1,
				Value: "",
			},
		},
		tokenTest{
			lex: "...",
			want: &language.Token{
				Kind:  language.TokenSpread,
				Start: 0,
				End:   3,
				Value: "",
			},
		},

		tokenTest{
			lex: ":",
			want: &language.Token{
				Kind:  language.TokenColon,
				Start: 0,
				End:   1,
				Value: "",
			},
		},
		tokenTest{
			lex: "=",
			want: &language.Token{
				Kind:  language.TokenEqual,
				Start: 0,
				End:   1,
				Value: "",
			},
		},
		tokenTest{
			lex: "@",
			want: &language.Token{
				Kind:  language.TokenAt,
				Start: 0,
				End:   1,
				Value: "",
			},
		},
		tokenTest{
			lex: "[",
			want: &language.Token{
				Kind:  language.TokenBracketLeft,
				Start: 0,
				End:   1,
				Value: "",
			},
		},
		tokenTest{
			lex: "]",
			want: &language.Token{
				Kind:  language.TokenBracketRight,
				Start: 0,
				End:   1,
				Value: "",
			},
		},
		tokenTest{
			lex: "{",
			want: &language.Token{
				Kind:  language.TokenBraceLeft,
				Start: 0,
				End:   1,
				Value: "",
			},
		},

		tokenTest{
			lex: "|",
			want: &language.Token{
				Kind:  language.TokenPipe,
				Start: 0,
				End:   1,
				Value: "",
			},
		},

		tokenTest{
			lex: "}",
			want: &language.Token{
				Kind:  language.TokenBraceRight,
				Start: 0,
				End:   1,
				Value: "",
			},
		},
	}

	for _, test := range set {
		checkToken(t, test.lex, test.want)
	}
}

func TestLexReportsUserfulUnknownCharacterError(t *testing.T) {
	set := [][]string{
		[]string{
			"..",
			"Syntax Error GraphQL request (1:1) Cannot parse the unexpected character \".\".",
		},
		[]string{
			"?",
			"Syntax Error GraphQL request (1:1) Cannot parse the unexpected character \"?\".",
		},
		[]string{
			"\u203B",
			"Syntax Error GraphQL request (1:1) " +
				"Cannot parse the unexpected character \"\\u203B\".",
		},
		[]string{
			"\u203B",
			"Syntax Error GraphQL request (1:1) " +
				"Cannot parse the unexpected character \"\\u203B\".",
		},
		[]string{
			"\u200b",
			"Syntax Error GraphQL request (1:1) " +
				"Cannot parse the unexpected character \"\\u200B\".",
		},
	}

	for _, test := range set {
		_, err := lexOne(test[0])
		testErr(t, err, test[1])
	}
}

func TestLexReportsUsefulInformationForDashesInNames(t *testing.T) {
	q := "a-b"
	s := language.NewSource(q)
	lexer := CreateLexer(s)

	firstToken, err := lexer.Advance()
	if err != nil {
		t.Fatal(err)
	}

	want := &language.Token{
		Kind:  language.TokenName,
		Start: 0,
		End:   1,
		Value: "a",
	}

	checkExistingToken(t, q, firstToken, want)

	_, err = lexer.Advance()
	testErr(t, err, "Syntax Error GraphQL request (1:3) Invalid number, expected digit but got: \"b\".")
}

func TestProducesDoubleLinkedListOfTokensIncludingComments(t *testing.T) {
	lexer := CreateLexer(language.NewSource(`{
  #comment
  field
}`))

	startToken := lexer.Token
	var endToken *language.Token
	var err error

	for {
		endToken, err = lexer.Advance()
		if err != nil {
			t.Fatal(err)
		}
		// Lexer advances over ignored comment tokens to make writing parsers
		// easier, but will include them in the linked list result.
		if endToken.Kind == language.TokenComment {
			t.Errorf("expected token kind %v got %v", language.TokenComment, endToken.Kind)
		}
		if endToken.Kind == language.TokenEOF {
			break
		}
	}

	if startToken.Prev != nil {
		t.Errorf("startToken.Prev should be nil but is %v", startToken.Prev)
	}

	if startToken.Next == nil {
		t.Errorf("startToken.Next should not be nil")
	}

	if endToken.Prev == nil {
		t.Errorf("endToken.Prev should not be nil")
	}

	if endToken.Next != nil {
		t.Errorf("endToken.Next should be nil but is %v", endToken.Next)
	}

	tokens := make([]*language.Token, 0)
	for tok := startToken; tok != nil; tok = tok.Next {
		if len(tokens) > 0 {
			// Tokens are double-linked, prev should point to last seen token.
			if tok.Prev != tokens[len(tokens)-1] {
				t.Error("tokens are not doubly linked, prev should point to last seen token")
			}

		}

		tokens = append(tokens, tok)
	}

	kinds := []language.TokenKind{
		language.TokenSOF,
		language.TokenBraceLeft,
		language.TokenComment,
		language.TokenName,
		language.TokenBraceRight,
		language.TokenEOF,
	}

	for idx, tok := range tokens {
		if tok.Kind != kinds[idx] {
			t.Errorf("expected kind %v got %v", kinds[idx], tok.Kind)
		}
	}
}
