package enviro

import (
	"fmt"
)

// ParseError is returned by [Parse] when the environment variables cannot be
// mapped into the target struct. The Cause field contains the underlying error
// from the env parser.
type ParseError struct {
	Cause error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("ParseError: failed to parse env: %v", e.Cause)
}
