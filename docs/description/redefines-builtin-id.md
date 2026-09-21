## redefines-builtin-id

_Go version_: 1.0; behavior changes in 1.21.

_Description_: Constant names like `false`, `true`, `nil`, function names like `append`, `make`,
and basic type names like `bool`, and `byte` are not reserved words of the language; therefore the can be redefined.
Even if possible, redefining these built in names can lead to bugs very difficult to detect.

_Configuration_: N/A

