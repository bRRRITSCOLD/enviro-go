package dotenv

import (
	"fmt"
	"strings"
)

// DotEnvNotExistError is returned when the .env or .env.example file cannot
// be found at the given path.
type DotEnvNotExistError struct {
	Path string
}

func (e *DotEnvNotExistError) Error() string {
	return fmt.Sprintf("DotEnvNotExistError: %s does not exist", e.Path)
}

// DotEnvStatError is returned when the file system check on the .env file
// fails for a reason other than the file not existing.
type DotEnvStatError struct {
	Path  string
	Cause error
}

func (e *DotEnvStatError) Error() string {
	return fmt.Sprintf("DotEnvStatError: cannot stat %s: %v", e.Path, e.Cause)
}

// DotEnvLoadError is returned when the .env file exists but cannot be parsed
// or loaded into the environment.
type DotEnvLoadError struct {
	Path  string
	Cause error
}

func (e *DotEnvLoadError) Error() string {
	return fmt.Sprintf("DotEnvLoadError: failed to load %s: %v", e.Path, e.Cause)
}

// DotEnvReadError is returned when the .env.example file exists but cannot be
// read or parsed.
type DotEnvReadError struct {
	Path  string
	Cause error
}

func (e *DotEnvReadError) Error() string {
	return fmt.Sprintf("DotEnvReadError: failed to read %s: %v", e.Path, e.Cause)
}

// DotEnvValidateError is returned when validation of the .env file fails.
type DotEnvValidateError struct {
	Path  string
	Cause error
}

func (e *DotEnvValidateError) Error() string {
	return fmt.Sprintf("env: cannot stat %s: %v", e.Path, e.Cause)
}

// DotEnvMissingKeysError is returned by [DotEnv.Validate] when one or more
// keys declared in the .env.example file are absent from the environment.
// The Keys field lists every missing key.
type DotEnvMissingKeysError struct {
	Keys []string
}

func (e *DotEnvMissingKeysError) Error() string {
	return fmt.Sprintf("DotEnvValidateError: missing required keys (declared in example ): %s", strings.Join(e.Keys, ", "))
}
