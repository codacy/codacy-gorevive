## unhandled-error

_Description_: This rule warns when errors returned by a function are not explicitly handled on the caller side.

_Configuration_: function names regexp patterns to ignore

Example:

```toml
[unhandled-error]
  arguments =["os\.(Create|WriteFile|Chmod)", "fmt\.Print", "myFunction", "net\..*", "bytes\.Buffer\.Write"]
```
