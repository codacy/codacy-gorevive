## superfluous-else

_Description_: To improve the readability of code, it is recommended to reduce the indentation as much as possible.
This rule highlights redundant _else-blocks_ that can be eliminated from the code.

_Configuration_: ([]string) rule flags. Available flags are:

- `preserveScope` (`preservescope`, `preserve-scope`): (string) do not suggest refactorings that would increase variable scope

Examples:

```toml
[rule.superfluous-else]
arguments = ["preserveScope"]
```

```toml
[rule.superfluous-else]
arguments = ["preserve-scope"]
```

