## import-alias-naming

_Description_: Aligns with Go's naming conventions, as outlined in the official
[blog post](https://go.dev/blog/package-names). It enforces clear and lowercase import alias names, echoing
the principles of good package naming. Users can follow these guidelines by default or define a custom regex rule.
Importantly, aliases with underscores ("_") are always allowed.

_Configuration_: (string) regular expression to match the aliases (default: ^[a-z][a-z0-9]{0,}$)

Example:

```toml
[rule.import-alias-naming]
  arguments =["^[a-z][a-z0-9]{0,}$"]
```

