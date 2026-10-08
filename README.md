# go-api-key

Building blocks for API key management in Go: a minimal validation interface and a local, file-backed implementation that loads `name<separator>key` pairs from a text file.

## Installation

```bash
go get github.com/ralvarezdev/go-api-key
```

## Usage

```go
import (
	"log/slog"
	"os"

	golocal "github.com/ralvarezdev/go-api-key/local"
)

logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

// The separator splits each line into "service name" and "API key"
service, err := golocal.NewService("=", logger)
if err != nil {
	panic(err)
}

// One "name=key" per line; blank lines and lines starting with # are skipped
if err = service.Load("api_keys.txt"); err != nil {
	panic(err)
}

if service.IsAPIKeyValid("some-key") {
	// accept the request
}
```

## API

- **`goapikey.BasicService`** — interface with `IsAPIKeyValid(apiKey string) bool`; also `ErrNilService`.
- **`local.NewService(nameSeparator, logger)`** — creates the service; returns `ErrEmptyNameSeparator` for an empty separator, and `logger` may be nil.
- **`local.Service`** — `Load(path)` reads the file line by line, ignoring blank, comment and invalid lines; `IsAPIKeyValid(apiKey)` checks whether the key was loaded.
- **`grpc`** — placeholder package that only contains `ErrNilGRPCInterceptions`.

## Project structure

```
interfaces.go, errors.go   BasicService interface and root errors
local/                     file-backed service, interface and errors
grpc/                      placeholder package
```

There are no tests.

## License

GNU General Public License v3.0. See [LICENSE](LICENSE).
