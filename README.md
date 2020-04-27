# Go secrets CLI

CLI for [go-secrets][1]

## Installation

- Download: `go get -u https://github.com/tecnologer/go-secrets-cli`
- Open source path: `cd $GOPATH/src/github.com/tecnologer/go-secrets-cli`
- Install: `go install`

## How to use

To display help type: `go-secrets-cli help`

```txt
* Set new secret:
        go-secrets-cli set -id <uuid> -key <string> -value <string>
* Get secret:
        go-secrets-cli get -id <uuid> [-key <string>]
* Remove secret:
        go-secrets-cli remove -id <uuid> -key <string>
```

[1]: https://github.com/Tecnologer/go-secrets
