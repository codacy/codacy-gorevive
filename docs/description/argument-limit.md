## argument-limit

_Description_: Warns when a function receives more parameters than the maximum set by the rule's configuration.
Enforcing a maximum number of parameters helps to keep the code readable and maintainable.

_Configuration_: (int) the maximum number of parameters allowed per function.

Example:

```toml
[rule.argument-limit]
arguments = [4]
```

