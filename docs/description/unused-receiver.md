## unused-receiver

_Description_: This rule warns on unused method receivers. Methods with unused receivers can be a symptom of an unfinished refactoring or a bug.

_Configuration_: Supports arguments with single of `map[string]any` with option `allowRegex` to provide additional to `_` mask to allowed unused receiver names, for example:

```toml
[rule.unused-receiver]
    Arguments = [ { allowRegex = "^_" } ]
```

allows any names started with `_`, not just `_` itself:

```go
func (_my *MyStruct) SomeMethod() {} // matches rule
```

