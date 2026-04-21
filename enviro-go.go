package enviro

import (
	"github.com/bRRRITSCOLD/enviro-go/dotenv"

	caarlos "github.com/caarlos0/env/v11"
)

// Options around environment and its definitions
type EnvConfig struct {
	// DotEnv is the path to the .env file to load (e.g. ".env").
	// When empty, no file is loaded.
	DotEnv string
	// DotEnvExample is the path to the .env.example file used to validate
	// that all required keys are present (e.g. ".env.example").
	// When empty, no validation is performed.
	DotEnvExample string
}

type Env[T any] struct {
	cfg T
}

// Config returns the parsed configuration struct.
func (e *Env[T]) Config() T {
	return e.cfg
}

// Parse parses/loads your environment
func Parse[T any](ec EnvConfig) (*Env[T], error) {
	de := &dotenv.DotEnv{}

	if ec.DotEnv != "" {
		de.Path = ec.DotEnv
		if err := de.Load(); err != nil {
			return nil, err
		}
	}

	if ec.DotEnvExample != "" {
		de.Example = ec.DotEnvExample
		if err := de.Validate(); err != nil {
			return nil, err
		}
	}

	cfg, err := caarlos.ParseAs[T]()
	if err != nil {
		return nil, &EnvParseError{Cause: err}
	}

	return &Env[T]{cfg: cfg}, nil
}
