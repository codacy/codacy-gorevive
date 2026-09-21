## cyclomatic

_Go version_: 1.0.

_Description_: [Cyclomatic complexity](https://en.wikipedia.org/wiki/Cyclomatic_complexity) is a measure of code complexity.
Enforcing a maximum complexity per function helps to keep code readable and maintainable.

_Configuration_: (int) the maximum function complexity. Default: `10`.

Configuration example:

```toml
[rule.cyclomatic]
arguments = [3]
```

