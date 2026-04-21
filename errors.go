package enviro

import (
	"fmt"
)

// EnvParseError is returned by [Parse] when the environment variables cannot
// be mapped into the target struct. The Cause field contains the underlying
// error from the env parser.
type EnvParseError struct {
	Cause error
}

func (e *EnvParseError) Error() string {
	return fmt.Sprintf("EnvParseError: failed to parse env: %v", e.Cause)
}
