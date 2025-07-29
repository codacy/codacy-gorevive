## unused-parameter

_Description_: This rule warns on unused parameters. Functions or methods with unused parameters can be a symptom of an unfinished refactoring or a bug.

_Configuration_: Supports arguments with single of `map[string]any` with option `allowRegex` (`allowregex`, `allow-regex`) to provide additional
to `_` mask to allowed unused parameter names.

Examples:

This allows any names starting with `_`, not just `_` itself:

```go
func SomeFunc(_someObj *MyStruct) {} // matches rule
```

```toml
[rule.unused-parameter]
arguments = [{ allowRegex = "^_" }]
```

```toml
[rule.unused-parameter]
arguments = [{ allow-regex = "^_" }]
```

