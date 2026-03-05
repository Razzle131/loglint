# loglint

# Table of Contents
* [About](#about)
* [Technologies](#technologies)
* [Clone repo](#clone-repo)
* [Quick start](#quick-start)
* [Configuration](#configuration)
* [Additions](#additions)

## About
Loglint is a static analysis tool for Go code that checks the correctness of logging messages. 

## Technologies
Project is created with:
* Golang v1.25.3
* [cleanenv](https://pkg.go.dev/github.com/ilyakaznacheev/cleanenv@v1.5.0)
* [testify](https://pkg.go.dev/github.com/stretchr/testify@v1.11.1) 

## Clone repo
```
git clone https://github.com/Razzle131/loglint.git
```

## Quick start
### Build from source
* Ensure that Go is installed on your machine and it`s version is equal or higther than 1.25.3:
  ```
  go version
  ```
* Install dependencies:
  ```
  go mod tidy
  ```
* Run unit tests:
  ```
  go test ./... -cover
  ```
  or
  ```
  make unit
  ```
* Run tests on testdata files:
  ```
  go run cmd/main.go ./logcheck/testdata/src/...
  ```
  or
  ```
  make test
  ```
### golangci-lint plugin
* Ensure that golangci-lint is installed on your machine:
  ```
  golangci-lint version
  ```
* Build plugin:
  ```
  golangci-lint custom -v
  ```
  or
  ```
  make plugin
  ```
* Test builded plugin on testdata:
  ```
  ./custom-gcl run -c ./.local.golangci.yml -v ./logcheck/testdata/src/slog ./logcheck/testdata/src/zap
  ```
  or
  ```
  make plugin-run
  ```
## Configuration
There are default settings for this linter, however you can toggle rules, write own sensetive patterns and functions in config.yaml file.     
To run linter with config file you will need to specify -config flag:
```
go run ./cmd/main.go -config /path/to/config.yaml
```
Linter searches config.yaml in working directory by default.

When using as golangci-plugin add to .golangci.yml (or use .local.golangci.yml) setting cfgPath, example:
```
linters:
  enable:
    - loglint
  settings:
    custom:
      loglint:
        type: "module"
        settings:
          cfgPath: /path/to/config.yaml
```
______

## Additions
### Sensetive data check
* Was made only for idents, because we have positive example with "token" in message in rule 4
### Parameter rules
* No examples were given, so linter checks each parameter according to the same rules
### Test on golang source code
* Linter found some issues:
  ```
  /usr/local/go/src/log/slog/logger_test.go:457:3: first letter must be in lower case: "Info(\"A\")"
  /usr/local/go/src/log/slog/logger_test.go:459:3: first letter must be in lower case: "Info(\"B\")"
  /usr/local/go/src/log/slog/example_logvaluer_secret_test.go:32:2: first letter must be in lower case: "logger.Info(\"permission granted\", \"user\", \"Perry\", \"token\", t)"
  ```
