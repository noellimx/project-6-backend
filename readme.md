# RTFM

# package name

gomoon

# HTTPS TLS Certificate

## Generate a self-signed certificate

`openssl req -nodes -x509 -newkey rsa:4096 -keyout server.key -out server.cert -sha256 -days 365`

This certificate is for testing as it is not signed by any CA.

## Configuration

Let's enforce no reading of variables from environment for now. All configurations (paths, variables etc) should be stored in a json file `config.json`. Path to this file is assigned to `$HOME/configFilePath/config.json` and will be parsed as a global configuration struct. The config is confidential and MUST NOT be commited into repository.

The json shape of global configuration can be found in package `config`

[ ] CI/CD : toggle development / production (deployment)

# Development

1. Clone the repository
2. `go install` will build a binary in `$GOPATH`
3. `$GOPATH/gomoon` execute binary

# Testing

`go test -v ./...` all files.
