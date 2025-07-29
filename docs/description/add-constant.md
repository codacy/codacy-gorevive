## add-constant

_Description_: Suggests using constant for [magic numbers](https://en.wikipedia.org/wiki/Magic_number_(programming)#Unnamed_numerical_constants)
and string literals.

_Configuration_:

- `maxLitCount` (`maxlitcount`, `max-lit-count`): (string) maximum number of instances of a string literal that are tolerated before warn.
- `allowStrs` (`allowstrs`, `allow-strs`): (string) comma-separated list of allowed string literals
- `allowInts` (`allowints`, `allow-ints`): (string) comma-separated list of allowed integers
- `allowFloats` (`allowfloats`, `allow-floats`): (string) comma-separated list of allowed floats
- `ignoreFuncs` (`ignorefuncs`, `ignore-funcs`): (string) comma-separated list of function names regexp patterns to exclude

Examples:

```toml
[rule.add-constant]
arguments = [
  { maxLitCount = "3", allowStrs = "\"\"", allowInts = "0,1,2", allowFloats = "0.0,0.,1.0,1.,2.0,2.", ignoreFuncs = "os\\.*,fmt\\.Println,make" },
]
```

```toml
[rule.add-constant]
arguments = [
  { max-lit-count = "3", allow-strs = "\"\"", allow-ints = "0,1,2", allow-floats = "0.0,0.,1.0,1.,2.0,2.", ignore-funcs = "os\\.*,fmt\\.Println,make" },
]
```

