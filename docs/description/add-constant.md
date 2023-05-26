## add-constant

_Description_: Suggests using constant for [magic numbers](https://en.wikipedia.org/wiki/Magic_number_(programming)#Unnamed_numerical_constants) and string literals.

_Configuration_:

* `maxLitCount` : (string) maximum number of instances of a string literal that are tolerated before warn.
* `allowStr`: (string) comma-separated list of allowed string literals
* `allowInts`: (string) comma-separated list of allowed integers
* `allowFloats`: (string) comma-separated list of allowed floats
* `ignoreFuncs`: (string) comma-separated list of function names regexp patterns to exclude

Example:

```toml
[rule.add-constant]
  arguments = [{ maxLitCount = "3", allowStrs = "\"\"", allowInts = "0,1,2", allowFloats = "0.0,0.,1.0,1.,2.0,2.", ignoreFuncs = "os\\.*,fmt\\.Println,make" }]
```

