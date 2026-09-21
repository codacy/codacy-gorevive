## cognitive-complexity

_Go version_: 1.0.

_Description_: [Cognitive complexity](https://www.sonarsource.com/docs/CognitiveComplexity.pdf) is a measure of how hard code is to understand.
While cyclomatic complexity is good to measure "testability" of the code,
cognitive complexity aims to provide a more precise measure of the difficulty of understanding the code.
Enforcing a maximum complexity per function helps to keep code readable and maintainable.

_Configuration_: (int) the maximum function complexity. Default: `7`.

Configuration example:

```toml
[rule.cognitive-complexity]
arguments = [7]
```

