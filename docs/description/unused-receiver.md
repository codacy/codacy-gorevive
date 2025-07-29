## unused-receiver

_Description_: This rule warns on unused method receivers. Methods with unused receivers can be a symptom of an unfinished refactoring or a bug.

_Configuration_:
Supports arguments with single of `map[string]any` with option `allowRegex` to provide additional to `_` mask to allowed unused receiver names.

Examples:

This allows any names starting with `_`, not just `_` itself:

```go
func (_my *MyStruct) SomeMethod() {} // matches rule
```

```toml
[rule.unused-receiver]
arguments = [{ allowRegex = "^_" }]
```

```toml
[rule.unused-receiver]
arguments = [{ allow-regex = "^_" }]
```

