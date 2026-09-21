## add-constant

_Go version_: 1.0.

_Description_: Suggests using constant for [magic numbers](https://en.wikipedia.org/wiki/Magic_number_(programming)#Unnamed_numerical_constants)
and string literals.

_Configuration_:

- `max-lit-count`: (string) maximum number of instances of a string literal that are tolerated before a warning is emitted.
- `allow-strs`: (string) comma-separated list of allowed string literals
- `allow-ints`: (string) comma-separated list of allowed integers
- `allow-floats`: (string) comma-separated list of allowed floats
- `ignore-funcs`: (string) comma-separated list of function names regexp patterns to exclude

Configuration example:

```toml
[rule.add-constant]
arguments = [
  { max-lit-count = "3", allow-strs = "\"\"", allow-ints = "0,1,2", allow-floats = "0.0,0.,1.0,1.,2.0,2.", ignore-funcs = "os\\.*,fmt\\.Println,make" },
]
```

