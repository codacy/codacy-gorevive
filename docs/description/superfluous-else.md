## superfluous-else

_Description_: To improve the readability of code, it is recommended to reduce the indentation as much as possible.
This rule highlights redundant _else-blocks_ that can be eliminated from the code.

_Configuration_: ([]string) rule flags. Available flags are:

* _preserveScope_: do not suggest refactorings that would increase variable scope

Example:

```toml
[rule.superfluous-else]
  arguments = ["preserveScope"]
```

