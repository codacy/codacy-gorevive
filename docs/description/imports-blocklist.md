## imports-blocklist

_Description_: Warns when importing block-listed packages.

_Configuration_: block-list of package names (or regular expression package names).

Example:

```toml
[imports-blocklist]
  arguments =["crypto/md5", "crypto/sha1", "crypto/**/pkix"]
```

