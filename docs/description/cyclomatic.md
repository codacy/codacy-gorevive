## cyclomatic

_Description_: [Cyclomatic complexity](https://en.wikipedia.org/wiki/Cyclomatic_complexity) is a measure of code complexity. Enforcing a maximum complexity per function helps to keep code readable and maintainable.

_Configuration_: (int) the maximum function complexity

Example:

```toml
[rule.cyclomatic]
  arguments = [3]
```

