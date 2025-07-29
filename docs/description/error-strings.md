## error-strings

_Description_: By convention, for better readability, error messages should not be capitalized or end with punctuation or a newline.
By default, the rule analyzes functions for creating errors from `fmt`, `errors`, and `github.com/pkg/errors`.
Optionally, the rule can be configured to analyze user functions that create errors.

More [information here](https://go.dev/wiki/CodeReviewComments#error-strings).

_Configuration_: ([]string) the list of additional error functions to check.
The format of values is `package.FunctionName`.

Example:

```toml
[rule.error-strings]
arguments = ["xerrors.Errorf"]
```

