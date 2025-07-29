## struct-tag

_Description_: Struct tags are not checked at compile time.
This rule spots errors in struct tags of the following types:
asn1, bson, datastore, default, json, mapstructure, properties, protobuf, required, toml, url, validate, xml, yaml.

_Configuration_: (optional) list of user defined options.

Example:
To accept the `inline` option in JSON tags (and `outline` and `gnu` in BSON tags) you must provide the following configuration

```toml
[rule.struct-tag]
arguments = ["json,inline", "bson,outline,gnu"]
```

