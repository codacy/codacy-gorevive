## max-control-nesting

_Go version_: 1.0.

_Description_: Warns if nesting level of control structures (`if-then-else`, `for`, `switch`) exceeds a given maximum.

_Configuration_: (int) maximum accepted nesting level of control structures. Default: `5`.

Configuration example:

```toml
[rule.max-control-nesting]
arguments = [3]
```

