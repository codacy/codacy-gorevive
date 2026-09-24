## unchecked-type-assertion

_Go version_: 1.0.

_Description_: This rule checks whether a type assertion result is checked (the `ok` value), preventing unexpected `panic`s.

_Configuration_: list of key-value-pair-map (`[]map[string]any`).

- `accept-ignored-assertion-result`: (bool) default `false`,
set it to `true` to accept ignored type assertion results like this:

```golang
foo, _ := bar(.*Baz).
//   ^
```

Configuration example:

```toml
[rule.unchecked-type-assertion]
arguments = [{ accept-ignored-assertion-result = true }]
```

