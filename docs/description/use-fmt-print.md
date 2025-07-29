## use-fmt-print

_Description_: This rule proposes to replace calls to built-in `print` and `println` with their equivalents from `fmt` standard package.

`print` and `println` built-in functions are not recommended for use-cases other than
[language boostraping and are not guaranteed to stay in the language](https://go.dev/ref/spec#Bootstrapping).

_Configuration_: N/A

