## function-result-limit

_Description_: Functions returning too many results can be hard to understand/use.

_Configuration_: (int) the maximum allowed return values

Example:

```toml
[rule.function-result-limit]
  arguments = [3]
```

