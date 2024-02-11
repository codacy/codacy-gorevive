## enforce-map-style

_Description_: This rule enforces consistent usage of `make(map[type]type)` or `map[type]type{}` for map initialization. It does not affect `make(map[type]type, size)` constructions as well as `map[type]type{k1: v1}`.

_Configuration_: (string) Specifies the enforced style for map initialization. The options are:
- "any": No enforcement (default).
- "make": Enforces the usage of `make(map[type]type)`.
- "literal": Enforces the usage of `map[type]type{}`.

Example:

```toml
[rule.enforce-map-style]
  arguments = ["make"]
```


