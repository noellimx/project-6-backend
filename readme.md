# Development

# RTFM

Each environment should have its own test and production configuration.

# package name

gomoon

# App Prerequisites

- go
- postgresql (go/gorm as ORM)

# githooks

git hooks are found in ./githooks.

## Example

plug hook on development environment

`cp ./dev-hooks/git-hooks/pre-commit .git/hooks/pre-commit`

## Github Action - CI/CD Part1 : Towards Deployment

# config.json [See](#configuration)

set secret variables for the gh environment via github repo settings. workflow will read from the settings and prepare config.json prior to running the program.

## mapping of names from config.json to github secrets

config.json key : Nesting = ".", Multi-word seperator: "\_", case: sensitive

=> github secret name : Nesting = "\_", Multi-word seperator: "", case: all uppercase

Example: config.db.my_secret => CONFIG_MYSECRET

Please note json keys are case-sensitive but github secret is case-insensitive

# Database Configuration

Using postgres with go/gorm. Application is initialized with gorm.AutoMigrate, that is a new development environment can start with a connection to an empty database.

## database_name

`gomoon` production \
`gomoontest` test

or any of your choice, just specify in json.

## One time setup for new environment

createdb <database_name>

# HTTPS TLS Certificate

## Generate a self-signed certificate

`openssl req -nodes -x509 -newkey rsa:4096 -keyout server.key -out server.cert -sha256 -days 365`

This certificate is fit for testing only as it is not signed by any CA.

## Configuration

No reading of variables from environment for now. All configurations (paths, variables to credentials, external services etc) should be stored in a json file `config.json`.

### Path

Path to this file is `$HOME/customkeystore/< "test" | "production" >/config.json` and will be parsed as a global configuration in the program. The config is confidential and MUST NOT be commited into repository. The binary will fatal if `config.json` cannot be detected. For example:

```Error Reading Config from path. open /home/ubuntu/customkeystore/production/config.json: no such file or directory```

### Syntax
The json shape of `config.json` can be found in package `config`

-----
[ ] CI/CD : toggle development / production (deployment)

# Development

1. Clone the repository
2. `go install` will build a binary in `$GOPATH`
3. `$GOPATH/bin/gomoon` execute binary

# Running package main

`go run .`

# Testing

`go test -v ./...` all go packages.

https://github.com/kaichung92/project-6-backend.git
