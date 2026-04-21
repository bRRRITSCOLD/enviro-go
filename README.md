# enviro-go

Type-safe environment variable parsing for Go with optional `.env` file loading and `.env.example` validation.

[![CI](https://github.com/bRRRITSCOLD/enviro-go/actions/workflows/ci.yml/badge.svg)](https://github.com/bRRRITSCOLD/enviro-go/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/bRRRITSCOLD/enviro-go.svg)](https://pkg.go.dev/github.com/bRRRITSCOLD/enviro-go)

## Install

```sh
go get github.com/bRRRITSCOLD/enviro-go
```

Requires Go 1.25+.

## Usage

Define a struct with `env` struct tags, then call `Parse`.

```go
type EnvironmentConfig struct {
    Host     string `env:"HOST,required"`
    Port     int    `env:"PORT" envDefault:"8080"`
    Debug    bool   `env:"DEBUG"`
}

env, err := enviro.Parse[EnvironmentConfig](enviro.EnvConfig{})
if err != nil {
    log.Fatal(err)
}

fmt.Println(env.Config().Host)
fmt.Println(env.Config().Port)
```

Struct tags are handled by [caarlos0/env](https://github.com/caarlos0/env) — see its documentation for the full tag syntax including required fields, defaults, slices, and custom parsers.

## Loading a .env file

Pass a path to `DotEnv` to load variables from a file before parsing. Variables already set in the environment take precedence.

```go
env, err := enviro.Parse[EnvironmentConfig](enviro.EnvConfig{
    DotEnv: ".env",
})
```

## Validating with .env.example

Pass a path to `DotEnvExample` to validate that every key declared in the example file is present in the environment. This is useful for catching missing variables at startup rather than at the point of use.

```go
env, err := enviro.Parse[EnvironmentConfig](enviro.EnvConfig{
    DotEnv:        ".env",
    DotEnvExample: ".env.example",
})
```

Given a `.env.example`:

```sh
HOST=
PORT=
```

`Parse` will return a `DotEnvMissingKeysError` if `HOST` or `PORT` are absent from the environment.

## Error handling

All errors are typed and can be inspected with `errors.As`.

| Error | Cause |
|---|---|
| `EnvParseError` | struct tag parsing failed |
| `DotEnvNotExistError` | `.env` or `.env.example` file not found |
| `DotEnvStatError` | file system error checking the file |
| `DotEnvLoadError` | `.env` file could not be parsed |
| `DotEnvReadError` | `.env.example` file could not be read |
| `DotEnvMissingKeysError` | one or more required keys are absent |

```go
var missing *dotenv.DotEnvMissingKeysError
if errors.As(err, &missing) {
    fmt.Println("missing keys:", missing.Keys)
}
```

## Development

```sh
# Install tools and git hooks
make setup

# Run tests
make test

# Run linter
make lint
```
