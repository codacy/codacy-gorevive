## enforce-slice-style

_Description_: This rule enforces consistent usage of `make([]type, 0)`, `[]type{}`, or `var []type` for slice initialization.
It does not affect `make([]type, non_zero_len, or_non_zero_cap)` constructions as well as `[]type{v1}`.
Nil slices are always permitted.

_Configuration_: (string) Specifies the enforced style for slice initialization. The options are:
- "any": No enforcement (default).
- "make": Enforces the usage of `make([]type, 0)`.
- "literal": Enforces the usage of `[]type{}`.
- "nil": Enforces the usage of `var []type`.

Example:

```toml
[rule.enforce-slice-style]
  arguments = ["make"]
```

