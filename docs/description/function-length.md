## function-length

_Description_: Functions too long (with many statements and/or lines) can be hard to understand.

_Configuration_: (int,int) the maximum allowed statements and lines. Must be non-negative integers. Set to 0 to disable the check

Example:

```toml
[rule.function-length]
arguments = [10, 0]
```

Will check for functions exceeding 10 statements and will not check the number of lines of functions

