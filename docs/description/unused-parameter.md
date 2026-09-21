## unused-parameter

_Go version_: 1.0.

_Description_: This rule warns on unused parameters. Functions or methods with unused parameters can be a symptom of an unfinished refactoring or a bug.

_Configuration_: Supports a single `map[string]any` argument with an `allow-regex` option to
specify additional allowed patterns for unused parameter names beyond the default `_`.

Configuration example:

This allows any names starting with `_`, not just `_` itself:

```go
func SomeFunc(_someObj *MyStruct) {} // matches rule
```

```toml
[rule.unused-parameter]
arguments = [{ allow-regex = "^_" }]
```

