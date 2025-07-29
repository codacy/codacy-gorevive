## unchecked-type-assertion

_Description_: This rule checks whether a type assertion result is checked (the `ok` value), preventing unexpected `panic`s.

_Configuration_: list of key-value-pair-map (`[]map[string]any`).

- `acceptIgnoredAssertionResult` (`acceptignoredassertionresult`, `accept-ignored-assertion-result`): (bool) default `false`,
set it to `true` to accept ignored type assertion results like this:

```golang
foo, _ := bar(.*Baz).
//   ^
```

Examples:

```toml
[rule.unchecked-type-assertion]
arguments = [{ acceptIgnoredAssertionResult = true }]
```

```toml
[rule.unchecked-type-assertion]
arguments = [{ accept-ignored-assertion-result = true }]
```

