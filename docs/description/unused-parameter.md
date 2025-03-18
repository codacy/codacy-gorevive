## unused-parameter

_Description_: This rule warns on unused parameters. Functions or methods with unused parameters can be a symptom of an unfinished refactoring or a bug.

_Configuration_: Supports arguments with single of `map[string]any` with option `allowRegex` to provide additional to `_` mask to allowed unused parameter names, for example:

```toml
[rule.unused-parameter]
    arguments = [{ allowRegex = "^_" }]
```

allows any names started with `_`, not just `_` itself:

```go
func SomeFunc(_someObj *MyStruct) {} // matches rule
```

