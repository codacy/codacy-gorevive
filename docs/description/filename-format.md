## filename-format

_Description_: enforces conventions on source file names. By default, the rule enforces filenames of the form `^[_A-Za-z0-9][_A-Za-z0-9-]*\.go$`.
Optionally, the rule can be configured to enforce other forms.

_Configuration_: (string) regular expression for source filenames.

Example:

```toml
[rule.filename-format]
arguments = ["^[_a-z][_a-z0-9]*\\.go$"]
```

