# Development

# RTFM

# package name

gomoon

# App Prerequisites

- go
- postgresql (go/gorm as ORM)
- Non-Windows OS

------------

Each environment (EC2, local, Github Runner etc) should have its own test and production configuration. In general, all environments will have the following steps before running the application.

# 1. Install [prerequisites](#app-prerequisites)


## 1.1 Go
Development will use golang and its utilities.

## 1.2 postgresql
For database service.


# 2. One time setup for new environment

## Configuration

No reading of variables from environment for now. All configurations (paths, variables to credentials, external services etc) should be stored in a json file `config.json`.

### Path

Path to this file is `$HOME/customkeystore/< "test" | "production" >/config.json` and will be parsed as a global configuration in the program. The config is confidential and MUST NOT be commited into repository.

The binary will fatal if `config.json` cannot be detected. Example log:

```Error Reading Config from path. open /home/ubuntu/customkeystore/production/config.json: no such file or directory```

### Syntax
The json shape and descriptions of `config.json` can be found in package `config`

#### How To

##### Configure Database

Using go/gorm as ORM in conjunction with postgresql. Application is initialized with gorm.AutoMigrate, thus application can run in new development environment with a connection to an empty database.

###### Create databases
###### <database_name>
`gomoon` for production \
`gomoontest` for test

or any of your choice, just specify in `config.json`.

`$ createdb <database_name>`


Database service should be ready for connection prior to running the application.

##### Generate HTTPS TLS Certificate

###### Self-signed certificate

`$ openssl req -nodes -x509 -newkey rsa:4096 -keyout server.key -out server.cert -sha256 -days 365`

This certificate is fit for testing only as it is not signed by any CA.

# 3. Download Repository and Test Source Code

## Download Example
`$ git clone .....` cloning from git-supported url.

## Test Example
`$ go test -v ./...` test all go packages.

pre-commit hook should also test all go packages.

# 4. Install Binary and Run

If you want to run the binary, it should be ready by now.

`$ go install` will build a binary in `$GOPATH`

`$ go run .` run the source code

OR

`$ $GOPATH/bin/gomoon` run binary

---------------------------

# CI/CD

## If you are changing git history

run the corresponding hooks.

### git hooks
git hooks are found in ./githooks.

### Example
plug hook on development environment
`$ cp ./dev-hooks/git-hooks/pre-commit .git/hooks/pre-commit`

## CI/CD Part 1 : Towards Integration (Github Action w/ Runner)

Workflow can be found in default github workflow folder.

# config.json

## See: [Configuration](#configuration)

Secret variables will be set as configuration values for the gh environment via github repo settings. Workflow will read from the settings and prepare `config.json` prior to running the program.

## Mapping of names from config.json to github secrets

config.json key : Nesting = `.`, Multi-word separator: `_`, case: sensitive

=> github secret name : Nesting = `_`, Multi-word separator: none, case: all uppercase

Example: config.db.my_secret => CONFIG_MYSECRET

Please note json keys are case-sensitive but github secret is case-insensitive

## CI/CD Part 2 : Towards Deployment (Github Action w/ EC2)

# config.json

## See: [One-time Configuration](#2-one-time-setup-for-new-environment)



