package dotenv

import (
	"fmt"
)

// NotExistError is returned when the .env file cannot be found at the given path.
type NotExistError struct {
	Path string
}

func (e *NotExistError) Error() string {
	return fmt.Sprintf("NotExistError: %s does not exist", e.Path)
}

// StatError is returned when the file system check on the .env file fails for
// a reason other than the file not existing.
type StatError struct {
	Path  string
	Cause error
}

func (e *StatError) Error() string {
	return fmt.Sprintf("StatError: cannot stat %s: %v", e.Path, e.Cause)
}

// LoadError is returned when the .env file exists but cannot be parsed or
// loaded into the environment.
type LoadError struct {
	Path  string
	Cause error
}

func (e *LoadError) Error() string {
	return fmt.Sprintf("LoadError: failed to load %s: %v", e.Path, e.Cause)
}
