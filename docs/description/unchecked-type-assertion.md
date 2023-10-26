## unchecked-type-assertion

_Description_: This rule checks whether a type assertion result is checked (the `ok` value), preventing unexpected `panic`s.

_Configuration_: list of key-value-pair-map (`[]map[string]any`).

- `acceptIgnoredAssertionResult` : (bool) default `false`, set it to `true` to accept ignored type assertion results like this:

```go
foo, _ := bar(.*Baz).
//   ^
```

Example:

```yaml
[rule.unchecked-type-assertion]
arguments = [{acceptIgnoredAssertionResult=true}]
```

