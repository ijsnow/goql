package errors

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/ijsnow/goql/internal/language"
)

// NewSyntaxError returns a GraphQLError representing a syntax error, containing useful
// descriptive information about the syntax error's position in the source.
func NewSyntaxError(
	source language.Source,
	position int,
	description string,
) error {
	location := language.GetLocation(source, position)

	return NewGraphQLError(
		fmt.Sprintf("Syntax Error GraphQL request (%v:%v) ", location.Line, location.Column)+
			description+"\n\n"+highlightSourceAtLocation(source, location),
		nil,
		&source,
		[]int{position},
		nil,
		nil,
	)
}

// highlightSourceAtLocation is a helpful description
// of the location of the error in the GraphQL Source document.
func highlightSourceAtLocation(source language.Source, location language.SourceLocation) string {
	line := location.Line
	prevLineNum := fmt.Sprintf("%d", line-1)
	lineNum := fmt.Sprintf("%d", line)
	nextLineNum := fmt.Sprintf("%d", line+1)
	padLen := utf8.RuneCountInString(nextLineNum)
	lines := regexp.MustCompile("\r\n|[\n\r]").Split(source.Body, -1)

	out := ""

	// prev line if exists
	if line >= 2 {
		out = lpad(padLen, prevLineNum) + ": " + lines[line-2] + "\n"
	}

	// line with error
	out += lpad(padLen, lineNum) +
		": " +
		lines[line-1] +
		"\n" +
		strings.Join(make([]string, location.Column+padLen+2), " ") +
		"^\n"

	// next line if exists
	if line < len(lines) {
		out += lpad(padLen, nextLineNum) + ": " + lines[line] + "\n"
	}

	return out
}

func lpad(num int, str string) string {
	return strings.Join(make([]string, num-len(str)+1), " ") + str
}
