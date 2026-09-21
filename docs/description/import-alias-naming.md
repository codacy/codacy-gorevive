## import-alias-naming

_Go version_: 1.0.

_Description_: Aligns with Go's naming conventions, as outlined in the official
[blog post](https://go.dev/blog/package-names). It enforces clear and lowercase import alias names, echoing
the principles of good package naming. Users can follow these guidelines by default or define a custom regex rule.
Importantly, aliases with underscores ("_") are always allowed.

_Configuration_ (1): (`string`) as plain string accepts allow regexp pattern for aliases (default: `^[a-z][a-z0-9]{0,}$`).

_Configuration_ (2): (`map[string]string`) as a map accepts two values:

- for a key `allow-regex` accepts allow regexp pattern
- for a key `deny-regex` deny regexp pattern

_Note_: If both `allow-regex` and `deny-regex` are provided, the alias must comply with both of them.
If none are given (i.e. an empty map), the default value `^[a-z][a-z0-9]{0,}$` for `allow-regex` is used.
Unknown keys will result in an error.

Configuration example (1):

```toml
[rule.import-alias-naming]
arguments = ["^[a-z][a-z0-9]{0,}$"]
```

Configuration example (2):

```toml
[rule.import-alias-naming]
arguments = [{ allow-regex = "^[a-z][a-z0-9]{0,}$", deny-regex = '^v\d+$' }]
```

