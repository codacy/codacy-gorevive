## imports-blacklist

_Description_: Warns when importing black-listed packages.

_Configuration_: black-list of package names (or regular expression package names).

Example:

```toml
[imports-blacklist]
  arguments =["crypto/md5", "crypto/sha1", "crypto/**/pkix"]
```

