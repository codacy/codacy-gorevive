## function-result-limit

_Go version_: 1.0.

_Description_: Specifies the maximum number of results a function can return.
Functions returning too many results can be hard to understand/use.

_Configuration_: (int) the maximum allowed return values. Default: `3`.

Configuration example:

```toml
[rule.function-result-limit]
arguments = [3]
```

