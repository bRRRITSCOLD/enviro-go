package dotenv

import (
	"os"

	"github.com/joho/godotenv"
)

type DotEnv struct {
	Path    string
	Example string
}

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
