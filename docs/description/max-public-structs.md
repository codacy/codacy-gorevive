## max-public-structs

_Description_: Packages declaring too many public structs can be hard to understand/use,
and could be a symptom of bad design.

This rule warns on files declaring more than a configured, maximum number of public structs.

_Configuration_: (int) the maximum allowed public structs

Example:

```toml
[rule.max-public-structs]
  arguments =[3]
```

