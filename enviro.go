// Package enviro provides type-safe environment variable parsing with optional
// .env file loading.
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
//		Path: ".env",
//	})
package enviro

import (
	"reflect"

	"github.com/bRRRITSCOLD/enviro-go/dotenv"
	caarlos "github.com/caarlos0/env/v11"
)

// ParserFunc is a function that parses a string into a value of a specific type.
// Use it to register custom type parsers via [EnvConfig.FuncMap].
type ParserFunc = caarlos.ParserFunc

// OnSetFn is a hook called whenever a field value is set during parsing.
// The tag is the env key, value is the resolved value, and isDefault indicates
// whether the value came from the envDefault tag rather than the environment.
type OnSetFn = caarlos.OnSetFn

// EnvConfig controls how [Parse] loads and parses the environment.
type EnvConfig struct {
	// Path is the path to a .env file to load before parsing (e.g. ".env").
	// When empty, no file is loaded.
	Path string

	// Environment overrides the source of environment variables. When set,
	// only the provided map is used for lookups instead of the process
	// environment. Use [github.com/caarlos0/env.ToMap] to convert
	// os.Environ() output into the expected format.
	Environment map[string]string

	// TagName overrides the struct tag used to identify environment variable
	// keys. Defaults to "env".
	TagName string

	// PrefixTagName overrides the struct tag used for env key prefixes on
	// nested structs. Defaults to "envPrefix".
	PrefixTagName string

	// DefaultValueTagName overrides the struct tag used for default values.
	// Defaults to "envDefault".
	DefaultValueTagName string

	// Prefix is prepended to every environment variable key before lookup.
	Prefix string

	// RequiredIfNoDef makes every field that has no envDefault tag required.
	RequiredIfNoDef bool

	// UseFieldNameByDefault uses the struct field name as the environment
	// variable key when no `env` tag is present on the field.
	UseFieldNameByDefault bool

	// SetDefaultsForZeroValuesOnly only applies envDefault values to fields
	// that are currently zero. Useful when mixing struct initialization values
	// with envDefault.
	SetDefaultsForZeroValuesOnly bool

	// OnSet is called after each field value is resolved.
	OnSet OnSetFn

	// FuncMap registers custom parser functions for specific types.
	FuncMap map[reflect.Type]ParserFunc
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
//  1. If [EnvConfig.Path] is set, the file is loaded into the environment.
//  2. The environment is parsed into T using the configured options.
//
// Any failure in the above steps returns a non-nil error.
func Parse[T any](ec EnvConfig) (*Env[T], error) {
	if ec.Path != "" {
		de := &dotenv.DotEnv{Path: ec.Path}
		if err := de.Load(); err != nil {
			return nil, err
		}
	}

	cfg, err := caarlos.ParseAsWithOptions[T](caarlos.Options{
		Environment:                  ec.Environment,
		TagName:                      ec.TagName,
		PrefixTagName:                ec.PrefixTagName,
		DefaultValueTagName:          ec.DefaultValueTagName,
		Prefix:                       ec.Prefix,
		RequiredIfNoDef:              ec.RequiredIfNoDef,
		UseFieldNameByDefault:        ec.UseFieldNameByDefault,
		SetDefaultsForZeroValuesOnly: ec.SetDefaultsForZeroValuesOnly,
		OnSet:                        ec.OnSet,
		FuncMap:                      ec.FuncMap,
	})
	if err != nil {
		return nil, &ParseError{Cause: err}
	}

	return &Env[T]{cfg: cfg}, nil
}
