package language

// TokenKind represents a GraphQL Token
type TokenKind string

// TokenKindToToken maps a TokenKind to the string that is the token
const (
	TokenSOF          TokenKind = "<SOF>"
	TokenEOF          TokenKind = "<EOF>"
	TokenBang         TokenKind = "!"
	TokenDollar       TokenKind = "$"
	TokenParenLeft    TokenKind = "("
	TokenParenRight   TokenKind = ")"
	TokenSpread       TokenKind = "..."
	TokenColon        TokenKind = ":"
	TokenEqual        TokenKind = "="
	TokenAt           TokenKind = "@"
	TokenBracketLeft  TokenKind = "["
	TokenBracketRight TokenKind = "]"
	TokenBraceLeft    TokenKind = "{"
	TokenBraceRight   TokenKind = "}"
	TokenPipe         TokenKind = "|"
	TokenName         TokenKind = "Name"
	TokenInt          TokenKind = "Int"
	TokenFloat        TokenKind = "Float"
	TokenString       TokenKind = "String"
	TokenComment      TokenKind = "Comment"
)

// Token represents a range of characters represented by a lexical token
// within a Source.
type Token struct {

	/**
	 * The kind of Token.
	 */
	Kind TokenKind `json:"kind"`

	/**
	 * The character offset at which this Node begins.
	 */
	Start int `json:"-"`

	/**
	 * The character offset at which this Node ends.
	 */
	End int `json:"-"`

	/**
	 * The 1-indexed line number on which this Token appears.
	 */
	Line int `json:"line"`

	/**
	 * The 1-indexed column number at which this Token begins.
	 */
	Column int `json:"column"`

	/**
	 * For non-punctuation tokens, represents the interpreted value of the token.
	 */
	Value string `json:"value"`

	/**
	 * Tokens exist as nodes in a double-linked-list amongst all tokens
	 * including ignored tokens. <SOF> is always the first node and <EOF>
	 * the last.
	 */
	Prev *Token `json:"-"`
	Next *Token `json:"-"`
}

// NewToken is a helper function for constructing the Token object.
func NewToken(
	kind TokenKind,
	start int,
	end int,
	line int,
	column int,
	prev *Token,
	value string,
) *Token {
	return &Token{
		Kind:   kind,
		Start:  start,
		End:    end,
		Line:   line,
		Column: column,
		Value:  value,
		Prev:   prev,
		Next:   nil,
	}
}
