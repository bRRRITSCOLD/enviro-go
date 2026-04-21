// Package dotenv provides .env file loading and .env.example validation.
package dotenv

import (
	"os"

	"github.com/joho/godotenv"
)

// DotEnv loads and validates .env files.
type DotEnv struct {
	// Path is the location of the .env file to load.
	Path string
	// Example is the location of the .env.example file used for validation.
	Example string
}

// Load reads the file at [DotEnv.Path] and sets its key-value pairs as
// environment variables. It returns an error if the file does not exist or
// cannot be parsed.
func (de *DotEnv) Load() error {
	_, err := os.Stat(de.Path)
	if os.IsNotExist(err) {
		return &DotEnvNotExistError{Path: de.Path}
	}

	if err != nil {
		return &DotEnvStatError{Path: de.Path, Cause: err}
	}

	if err := godotenv.Load(de.Path); err != nil {
		return &DotEnvLoadError{Path: de.Path, Cause: err}
	}

	return nil
}

// Validate reads the file at [DotEnv.Example] and checks that every key
// declared in it is present in the current environment. It returns a
// [DotEnvMissingKeysError] listing any keys that are absent.
func (de *DotEnv) Validate() error {
	_, err := os.Stat(de.Example)
	if os.IsNotExist(err) {
		return &DotEnvNotExistError{Path: de.Path}
	}
	if err != nil {
		return &DotEnvStatError{Path: de.Path, Cause: err}
	}

	example, err := godotenv.Read(de.Example)
	if err != nil {
		return &DotEnvReadError{Path: de.Path, Cause: err}
	}

	var missing []string
	for key := range example {
		if _, ok := os.LookupEnv(key); !ok {
			missing = append(missing, key)
		}
	}

	if len(missing) > 0 {
		return &DotEnvMissingKeysError{Keys: missing}
	}

	return nil
}
