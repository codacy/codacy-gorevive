## import-alias-naming

_Description_: Aligns with Go's naming conventions, as outlined in the official
[blog post](https://go.dev/blog/package-names). It enforces clear and lowercase import alias names, echoing
the principles of good package naming. Users can follow these guidelines by default or define a custom regex rule.
Importantly, aliases with underscores ("_") are always allowed.

_Configuration_ (1): (string) as plain string accepts allow regexp pattern for aliases (default: ^[a-z][a-z0-9]{0,}$).

_Configuration_ (2): (map[string]string) as a map accepts two values:
* for a key "allowRegex" accepts allow regexp pattern
* for a key "denyRegex deny regexp pattern

_Note_: If both `allowRegex` and `denyRegex` are provided, the alias must comply with both of them.
If none are given (i.e. an empty map), the default value ^[a-z][a-z0-9]{0,}$ for allowRegex is used.
Unknown keys will result in an error.

Example (1):

```toml
[rule.import-alias-naming]
  arguments = ["^[a-z][a-z0-9]{0,}$"]
```

Example (2):

```toml
[rule.import-alias-naming]
  arguments = [ { allowRegex = "^[a-z][a-z0-9]{0,}$", denyRegex = '^v\d+$' } ]
```

