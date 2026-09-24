## comment-spacings

_Go version_: 1.0.

_Description_: Warns on malformed comments.
Spots comments of the form:

```go
//This is a malformed comment: no space between // and the start of the sentence
```

_Configuration_: ([]string) list of exceptions. For example, to accept comments of the form

```go
//mypragma: activate something
//+optional
```

You need to add both `"mypragma:"` and `"+optional"` in the configuration

The following comment prefixes are allowed by default:

- `//#nosec` — [gosec](https://github.com/securego/gosec) security scanner directive
- Go directive comments matching `//[a-z0-9]+:[a-z0-9]` (e.g. `//nolint:linter`, `//go:generate`, `//revive:disable:rule`),
as well as comments starting with "//line ", "//extern ", and "//export "

Configuration example:

```toml
[rule.comment-spacings]
arguments = ["mypragma:", "+optional"]
```

