## function-length

_Go version_: 1.0.

_Description_: Functions too long (with many statements and/or lines) can be hard to understand.

_Configuration_: (int, int) the maximum allowed statements and lines.
Set a value to `0` to disable that specific check; if both values are `0`, the rule is disabled. Default: `50`, `75`.

Configuration example:

```toml
[rule.function-length]
arguments = [10, 0]
```

Will check for functions exceeding 10 statements and will not check the number of lines of functions

