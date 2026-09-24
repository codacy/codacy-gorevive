## superfluous-else

_Go version_: 1.0.

_Description_: To improve the readability of code, it is recommended to reduce the indentation as much as possible.
This rule highlights redundant _else-blocks_ that can be eliminated when the preceding `if`-block deviates control flow,
for example ending with a `break`, `continue`, `goto`, `panic` or `os.Exit` call (the `return` case is handled by [indent-error-flow](#indent-error-flow)).

