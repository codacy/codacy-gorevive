## receiver-naming

_Description_: By convention, receiver names in a method should reflect their identity.
For example, if the receiver is of type `Parts`, `p` is an adequate name for it.
Contrary to other languages, it is not idiomatic to name receivers as `this` or `self`.

_Configuration_: (optional) list of key-value-pair-map (`[]map[string]any`).

- `maxLength` (`maxlength`, `max-length`): (int) max length of receiver name

Examples:

```toml
[rule.receiver-naming]
arguments = [{ maxLength = 2 }]
```

```toml
[rule.receiver-naming]
arguments = [{ max-length = 2 }]
```

