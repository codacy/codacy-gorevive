## struct-tag

_Description_: Struct tags are not checked at compile time.
This rule, checks and warns if it finds errors in common struct tags types like: asn1, default, json, protobuf, xml, yaml.

_Configuration_: (optional) list of user defined options.

Example:
To accept the `inline` option in JSON tags (and `outline` and `gnu` in BSON tags) you must provide the following configuration

```toml
[rule.struct-tag]
  arguments = ["json,inline","bson,outline,gnu"]
```


