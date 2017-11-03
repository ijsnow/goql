package errors

import (
	"errors"

	"github.com/ijsnow/goql/internal/language"
)

// NewFormatError takes a GraphQLError and formats it according to the rules described by the
// Response Format, Errors section of the GraphQL Specification.
func NewFormatError(gqlerr *GraphQLError) (*GraphQLFormattedError, error) {
	if gqlerr == nil {
		return nil, errors.New("cannot format nil error")
	}

	return &GraphQLFormattedError{
		Message:   gqlerr.Message,
		Locations: gqlerr.Locations,
		Path:      gqlerr.Path,
	}, nil
}

// GraphQLFormattedError is an error formatted for marshalling
type GraphQLFormattedError struct {
	Message   string                    `json:"message"`
	Locations []language.SourceLocation `json:"locations"`
	Path      []string                  `json:"path"`
}
