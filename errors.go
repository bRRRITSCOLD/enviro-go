package enviro

import (
	"fmt"
)

type EnvParseError struct {
	Cause error
}

func (e *EnvParseError) Error() string {
	return fmt.Sprintf("EnvParseError: failed to parse env: %v", e.Cause)
}
