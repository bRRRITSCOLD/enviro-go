// Package dotenv provides .env file loading.
package dotenv

import (
	"os"

	"github.com/joho/godotenv"
)

// DotEnv loads a .env file into the environment.
type DotEnv struct {
	// Path is the location of the .env file to load.
	Path string
}

// Load reads the file at [DotEnv.Path] and sets its key-value pairs as
// environment variables. It returns an error if the file does not exist or
// cannot be parsed.
func (de *DotEnv) Load() error {
	_, err := os.Stat(de.Path)
	if os.IsNotExist(err) {
		return &NotExistError{Path: de.Path}
	}

	if err != nil {
		return &StatError{Path: de.Path, Cause: err}
	}

	if err := godotenv.Load(de.Path); err != nil {
		return &LoadError{Path: de.Path, Cause: err}
	}

	return nil
}
