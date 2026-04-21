package dotenv

import (
	"fmt"
	"strings"
)

type DotEnvNotExistError struct {
	Path string
}

func (e *DotEnvNotExistError) Error() string {
	return fmt.Sprintf("DotEnvNotExistError: %s does not exist", e.Path)
}

type DotEnvStatError struct {
	Path  string
	Cause error
}

func (e *DotEnvStatError) Error() string {
	return fmt.Sprintf("DotEnvStatError: cannot stat %s: %v", e.Path, e.Cause)
}

type DotEnvLoadError struct {
	Path  string
	Cause error
}

func (e *DotEnvLoadError) Error() string {
	return fmt.Sprintf("DotEnvLoadError: failed to load %s: %v", e.Path, e.Cause)
}

type DotEnvReadError struct {
	Path  string
	Cause error
}

func (e *DotEnvReadError) Error() string {
	return fmt.Sprintf("DotEnvReadError: failed to read %s: %v", e.Path, e.Cause)
}

type DotEnvValidateError struct {
	Path  string
	Cause error
}

func (e *DotEnvValidateError) Error() string {
	return fmt.Sprintf("env: cannot stat %s: %v", e.Path, e.Cause)
}

type DotEnvMissingKeysError struct {
	Keys []string
}

func (e *DotEnvMissingKeysError) Error() string {
	return fmt.Sprintf("DotEnvValidateError: missing required keys (declared in example ): %s", strings.Join(e.Keys, ", "))
}
