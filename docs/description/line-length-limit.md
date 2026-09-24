## line-length-limit

_Go version_: 1.0.

_Description_: Warns in the presence of code lines longer than a configured maximum.

_Configuration_: the rule accepts either an integer (the maximum line length) or a map of options:

- `max`: (int) maximum line length in characters. Default: `80`.
- `excludes`: ([]string) list of regular expressions; a line matching any of them is not checked. Default: none.

The integer form is a shorthand for setting only `max`.

Configuration examples:

```toml
[rule.line-length-limit]
arguments = [80]
```

```toml
[rule.line-length-limit]
arguments = [{ max = 80, excludes = ['^\s*//go:generate ', 'https?://'] }]
```

