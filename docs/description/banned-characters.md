## banned-characters

_Description_: Checks given banned characters in identifiers(func, var, const). Comments are not checked.

_Configuration_: This rule requires a slice of strings, the characters to ban.

Configuration example:

```toml
[rule.banned-characters]
arguments = ["Ω", "Σ", "σ"]
```

