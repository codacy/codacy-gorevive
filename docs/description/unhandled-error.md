## unhandled-error

_Go version_: 1.0.

**_Typed_**

_Description_: This rule warns when errors returned by a function are not explicitly handled on the caller side.

_Configuration_: function names regexp patterns to ignore

Configuration example:

```toml
[rule.unhandled-error]
arguments = [
  '^os\.(CreateTemp|WriteFile|Chmod)$',
  '^fmt\.Print',
  'myFunction',
  '^net\.',
  '^(bytes\.Buffer|string\.Writer)\.Write(Byte|Rune|String)?$',
]
```

