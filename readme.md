
# Development

# RTFM

Each environment should have its own test and production configuration.


# githooks

git hooks are found in ./githooks.

## Example
plug hook on development environment

```cp ./githooks/pre-commit .git/hooks/pre-commit```


# package name

gomoon

# Database Configuration

Using postgres


## database_name

`gomoon` production
`gomoontest` test

## One time setup for new environment

createdb <database_name>

# HTTPS TLS Certificate

## Generate a self-signed certificate

`openssl req -nodes -x509 -newkey rsa:4096 -keyout server.key -out server.cert -sha256 -days 365`

This certificate is fit for testing only as it is not signed by any CA.

## Configuration

Let's enforce no reading of variables from environment for now. All configurations (paths, variables etc) should be stored in a json file `config.json`. Path to this file is assigned to `$HOME/customkeystore/< "test" | "production" >/config.json` and will be parsed as a global configuration struct. The config is confidential and MUST NOT be commited into repository.

The json shape of global configuration can be found in package `config`

[ ] CI/CD : toggle development / production (deployment)

# Development

1. Clone the repository
2. `go install` will build a binary in `$GOPATH`
3. `$GOPATH/gomoon` execute binary

# Running package main

`go run .`

# Testing

`go test -v ./...` all files.

https://github.com/kaichung92/project-6-backend.git
