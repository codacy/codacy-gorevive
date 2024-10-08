## max-control-nesting

_Description_: Warns if nesting level of control structures (`if-then-else`, `for`, `switch`) exceeds a given maximum.

_Configuration_: (int) maximum accepted nesting level of control structures (defaults to 5)

Example:

```toml
[rule.max-control-nesting]
  arguments = [3]
```

