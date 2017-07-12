package language

import "regexp"

// SourceLocation represents a location in a Source.
type SourceLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

// GetLocation takes a Source and a UTF-8 character offset, and returns the corresponding
// line and column as a SourceLocation.
func GetLocation(source Source, position int) SourceLocation {
	lineRegexp := regexp.MustCompile("\r\n|[\n\r]")

	line := 1
	column := position + 1

	ms := lineRegexp.FindAllStringIndex(source.Body, -1)
	for _, match := range ms {
		if match[0] < position {
			line++
			column = position + 1 - (match[0] + (match[1] - match[0]))
		} else {
			break
		}
	}

	return SourceLocation{
		Line:   line,
		Column: column,
	}
}
