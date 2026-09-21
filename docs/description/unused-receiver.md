## unused-receiver

_Go version_: 1.0.

_Description_: This rule warns on unused method receivers. Methods with unused receivers can be a symptom of an unfinished refactoring or a bug.

_Configuration_:
Supports a single `map[string]any` argument with an `allow-regex` option to specify additional allowed unused receiver name patterns beyond `_`.

Configuration example:

This allows any names starting with `_`, not just `_` itself:

```go
func (_my *MyStruct) SomeMethod() {} // matches rule
```

```toml
[rule.unused-receiver]
arguments = [{ allow-regex = "^_" }]
```

