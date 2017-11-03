package errors

import "github.com/ijsnow/goql/internal/language"

// NewLocatedError , given an arbitrary Error, presumably thrown while attempting to execute a
// GraphQL operation, produce a new GraphQLError aware of the location in the
// document responsible for the original Error.
func NewLocatedError(
	originalError *GraphQLError,
	nodes []language.ASTNode,
	path []string,
) error {
	// Note: this uses a brand-check to support GraphQL errors originating from
	// other contexts.
	if originalError != nil && len(originalError.Path) > 0 {
		return originalError
	}

	var message string
	if originalError != nil {
		message = originalError.Message
	} else {
		message = "An unknown error occurred."
	}

	return NewGraphQLError(
		message,
		originalError.Nodes,
		originalError.Source,
		originalError.Positions,
		path,
		originalError,
	)
}
