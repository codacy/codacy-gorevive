## exported

_Go version_: 1.0.

**_Ported from golint_**

_Description_: Exported function and methods should have comments. This warns on undocumented exported functions and methods.

More [information here](https://go.dev/wiki/CodeReviewComments#doc-comments).

_Configuration_: ([]string) rule flags.
Please notice that without configuration, the default behavior of the rule is that of its `golint` counterpart.
Available flags are:

- `check-private-receivers` enables checking public methods of private types
- `disable-stuttering-check` disables checking for method names that stutter with the package name
  (i.e. avoid failure messages of the form _type name will be used as x.XY by other packages, and that stutters; consider calling this Y_)
- `say-repetitive-instead-of-stutters` replaces the use of the term _stutters_ by _repetitive_ in failure messages
- `check-public-interface` enables checking public method definitions in public interface types
- `disable-checks-on-constants` disables all checks on constant declarations
- `disable-checks-on-functions` disables all checks on function declarations
- `disable-checks-on-methods` disables all checks on method declarations
- `disable-checks-on-types` disables all checks on type declarations
- `disable-checks-on-variables` disables all checks on variable declarations

Configuration example:

```toml
[rule.exported]
arguments = [
  "check-private-receivers",
  "disable-stuttering-check",
  "check-public-interface",
  "disable-checks-on-functions",
]
```

