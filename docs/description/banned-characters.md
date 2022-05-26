## banned-characters

_Description_: Checks given banned characters in identifiers(func, var, const). Comments are not checked.

_Configuration_: This rule accepts a slice of strings, the characters to ban.

Example:

```toml
[rule.banned-characters]
  arguments =["Ω","Σ","σ"]
```

