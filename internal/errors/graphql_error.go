package errors

import (
	"github.com/ijsnow/goql/internal/language"
)

// GraphQLError describes an Error found during the parse, validate, or
// execute phases of performing a GraphQL operation. In addition to a message
// and stack trace, it also includes information about the locations in a
// GraphQL document and/or execution result that correspond to the Error.
type GraphQLError struct {

	/* Message describing the Error for debugging purposes.
	 * Enumerable, and appears in the result of JSON.stringify().
	 */
	Message string `json:"message"`

	/* Locations is an array of { line, column } locations within the source GraphQL document
	 * which correspond to this error.
	 *
	 * Errors during validation often contain multiple locations, for example to
	 * point out two things with the same name. Errors during execution include a
	 * single location, the field which produced the error.
	 *
	 * Enumerable, and appears in the result of JSON.stringify().
	 */
	Locations []language.SourceLocation `json:"locations"`

	/* Path is an array describing the JSON-path into the execution response which
	 * corresponds to this error. Only included for errors during execution.
	 *
	 * Enumerable, and appears in the result of JSON.stringify().
	 */
	Path []string `json:"path"`

	/**
	 * An array of GraphQL AST Nodes corresponding to this error.
	 */
	Nodes language.NodeList `json:"-"`

	/**
	 * The source GraphQL document corresponding to this error.
	 */
	Source *language.Source `json:"-"`

	/* Positions is an array of character offsets within the source GraphQL document
	 * which correspond to this error.
	 */
	Positions []int `json:"_"`

	/* OriginalError the original error thrown from a field resolver during execution.
	 */
	OriginalError error `json:"-"`
}

func (e GraphQLError) Error() string {
	return e.Message
}

// NewGraphQLError creates a new GraphQLError
func NewGraphQLError(
	message string,
	nodes language.NodeList,
	source *language.Source,
	positions []int,
	path []string,
	originalError error,
) error {
	// Compute locations in the source for the given nodes/positions.
	_source := source
	if _source == nil && nodes != nil && len(nodes) > 0 {
		loc := nodes[0].GetLoc()

		if loc != nil {
			_source = &loc.Source
		}
	}

	_positions := positions
	if _positions == nil && nodes != nil {
		_positions = nodes.GetStarts()
	}

	if _positions != nil && len(_positions) == 0 {
		_positions = nil
	}

	var _locations []language.SourceLocation
	if _source != nil && _positions != nil {
		for _, pos := range _positions {
			_locations = append(_locations, language.GetLocation(*_source, pos))
		}
	}

	gqlerr := GraphQLError{
		Message:       message,
		Locations:     _locations,
		Path:          path,
		Nodes:         nodes,
		Source:        _source,
		Positions:     _positions,
		OriginalError: originalError,
	}

	// TODO: Equivalent of this in go?
	// Include (non-enumerable) stack trace.
	//   if (originalError && originalError.stack) {
	//     Object.defineProperty(this, 'stack', {
	//       value: originalError.stack,
	//       writable: true,
	//       configurable: true
	//     });
	//   } else if (Error.captureStackTrace) {
	//     Error.captureStackTrace(this, GraphQLError);
	//   } else {
	//     Object.defineProperty(this, 'stack', {
	//       value: Error().stack,
	//       writable: true,
	//       configurable: true
	//     });
	//   }

	return gqlerr
}
