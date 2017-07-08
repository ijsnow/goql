package language

// Source is a representation of source input to GraphQL
type Source struct {
	Body string
}

// NewSource creates a new Source struct
func NewSource(body string) Source {
	return Source{
		Body: body,
	}
}
