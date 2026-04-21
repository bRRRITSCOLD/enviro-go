// Package enviro provides type-safe environment variable parsing with optional
// .env file loading and .env.example validation.
//
// Use [Parse] to load environment variables into a typed struct. Struct fields
// are mapped to environment variables using the `env` struct tag, powered by
// [github.com/caarlos0/env].
//
// # Basic usage
//
//	type Environment struct {
//		Host string `env:"HOST,required"`
//		Port int    `env:"PORT" envDefault:"8080"`
//	}
//
//	env, err := enviro.Parse[Environment](enviro.EnvConfig{})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(env.Config().Host)
//
// # Loading a .env file
//
//	env, err := enviro.Parse[Environment](enviro.EnvConfig{
//		DotEnv: ".env",
//	})
//
// # Validating against a .env.example
//
//	env, err := enviro.Parse[Environment](enviro.EnvConfig{
//		DotEnv:        ".env",
//		DotEnvExample: ".env.example",
//	})
package enviro

import (
	"github.com/bRRRITSCOLD/enviro-go/dotenv"

	caarlos "github.com/caarlos0/env/v11"
)

// EnvConfig controls how [Parse] loads and validates the environment.
type EnvConfig struct {
	// DotEnv is the path to a .env file to load before parsing (e.g. ".env").
	// When empty, no file is loaded.
	DotEnv string
	// DotEnvExample is the path to a .env.example file whose keys are used to
	// validate that all required variables are present in the environment.
	// When empty, no validation is performed.
	DotEnvExample string
}

// Env holds the parsed configuration of type T.
type Env[T any] struct {
	cfg T
}

// Config returns the parsed configuration struct.
func (e *Env[T]) Config() T {
	return e.cfg
}

// Parse loads the environment into a value of type T and returns it wrapped in
// an [Env]. The type parameter T must be a struct whose fields are annotated
// with `env` struct tags.
//
// Parse proceeds in order:
//  1. If [EnvConfig.DotEnv] is set, the file is loaded into the environment.
//  2. If [EnvConfig.DotEnvExample] is set, every key declared in that file is
//     checked for presence in the environment.
//  3. The environment is parsed into T using the `env` struct tags.
//
// Any failure in the above steps returns a non-nil error.
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
