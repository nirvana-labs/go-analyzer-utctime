# golangci-lint-utctime

A custom linter for [golangci-lint](https://golangci-lint.run/) that ensures all `time.Now()` calls are followed by `.UTC()`.

## Description

This linter helps prevent timezone-related bugs by ensuring that all `time.Now()` calls are immediately followed by `.UTC()`. This is particularly useful in applications where consistent timezone handling is critical.

## Installation

```bash
go install github.com/yourusername/golangci-lint-utctime@latest
```

## Usage

Add the linter to your `.golangci.yml` configuration:

```yaml
linters:
  enable:
    - utctime

linters-settings:
  custom:
    utctime:
      path: github.com/nirvana-labs/golangci-lint-utctime
```

## Examples

```go
// Bad:
t := time.Now() // Will trigger a linter error

// Good:
t := time.Now().UTC()
```

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
