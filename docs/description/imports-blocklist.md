## imports-blocklist

_Go version_: 1.0.

_Description_: Warns when importing block-listed packages.

_Configuration_: block-list of package names (or regular expression package names).

Configuration example:

```toml
[rule.imports-blocklist]
arguments = ["crypto/md5", "crypto/sha1", "crypto/**/pkix"]
```

