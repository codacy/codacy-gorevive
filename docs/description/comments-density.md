## comments-density

_Go version_: 1.0.

_Description_: Spots files not respecting a minimum value for the [_comments lines density_](https://docs.sonarsource.com/sonarqube/latest/user-guide/metric-definitions/)
metric = _comment lines / (lines of code + comment lines) * 100_

_Configuration_: (int) the minimum expected comments lines density. Default: `0`.

Configuration example:

```toml
[rule.comments-density]
arguments = [15]
```

