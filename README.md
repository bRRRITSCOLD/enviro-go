# enviro-go

Type-safe environment variable parsing for Go with optional `.env` file loading.

[![CI](https://github.com/bRRRITSCOLD/enviro-go/actions/workflows/ci.yml/badge.svg)](https://github.com/bRRRITSCOLD/enviro-go/actions/workflows/ci.yml)
[![Coverage](https://codecov.io/gh/bRRRITSCOLD/enviro-go/branch/main/graph/badge.svg)](https://codecov.io/gh/bRRRITSCOLD/enviro-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/bRRRITSCOLD/enviro-go.svg)](https://pkg.go.dev/github.com/bRRRITSCOLD/enviro-go)

## Install

```sh
go get github.com/bRRRITSCOLD/enviro-go
```

Requires Go 1.25+.

## Usage

Define a struct with `env` struct tags, then call `Parse`.

```go
type Environment struct {
    Host  string `env:"HOST,required"`
    Port  int    `env:"PORT" envDefault:"8080"`
    Debug bool   `env:"DEBUG"`
}

cfg, err := enviro.Parse[Environment](enviro.EnvConfig{})
if err != nil {
    log.Fatal(err)
}

fmt.Println(cfg.Config().Host)
fmt.Println(cfg.Config().Port)
```

Struct tags are handled by [caarlos0/env](https://github.com/caarlos0/env) — see its documentation for the full tag syntax including required fields, defaults, slices, and custom parsers.

## Loading a .env file

Pass a path to `DotEnv` to load variables from a file before parsing. Variables already set in the environment take precedence.

```go
cfg, err := enviro.Parse[Environment](enviro.EnvConfig{
    DotEnv: ".env",
})
```

> It is recommended to include a `.env.example` file in your repository documenting the expected variables. This gives developers a template to copy when setting up the project locally.

## Error handling

All errors are typed and can be inspected with `errors.As`.

| Error | Cause |
|---|---|
| `enviro.ParseError` | struct tag parsing failed |
| `dotenv.NotExistError` | `.env` file not found |
| `dotenv.StatError` | file system error checking the `.env` file |
| `dotenv.LoadError` | `.env` file could not be parsed |

```go
var notExist *dotenv.NotExistError
if errors.As(err, &notExist) {
    fmt.Println("missing .env file:", notExist.Path)
}
```

## Development

```sh
# Install tools and git hooks
make setup

# Run tests Linux/Unix
make test
# or
go run ./scripts/test/main.go
# or
go test -v -race -count=1 ./...

# Run tests Windows
make test
# or
go run ./scripts/test/main.go
# or
go test -v -count=1 ./...

# Run linter
make lint
```
