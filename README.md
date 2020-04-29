# Go secrets CLI

CLI for [go-secrets][1]

## Installation

- Download: `go get -u github.com/tecnologer/go-secrets-cli`
- Open source path: `cd $GOPATH/src/github.com/tecnologer/go-secrets-cli`
- Install: `go install`

## How to use

To display help type: `go-secrets-cli help`

```txt
* Set new secret:
        go-secrets-cli set [-id <uuid>] -key <string> -value <string>
* Get secret:
        go-secrets-cli get [-id <uuid>] [-key <string>]
* Remove secret:
        go-secrets-cli remove [-id <uuid>] -key <string>
* Init secret:
        go-secrets-cli init
```

### Tricks

- Omit `-id`
  > If the secret is initialized in the current folder, you can ommit the `-id` on each call.
  >
  > - `go-secrets-cli init`
  > - `go-secrets-cli get -key username`
- Create a group
  > To create a group, use a point between the group name and the key.
  > I.e: `go-secrets-cli set -key SQL.username -val tecnologer`

## ToDo

- [ ] Get all keys in a group

[1]: https://github.com/Tecnologer/go-secrets
