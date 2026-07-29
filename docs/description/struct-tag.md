## struct-tag

_Description_: The rule spots errors in struct tags.
This is useful because struct tags are not checked at compile time.

The list of [supported tags](https://go.dev/wiki/Well-known-struct-tags):

| Tag           | Documentation                                                            |
| ------------- | ------------------------------------------------------------------------ |
| `asn1`         | <https://pkg.go.dev/encoding/asn1>                                      |
| `bson`         | <https://pkg.go.dev/go.mongodb.org/mongo-driver/bson>                   |
| `cbor`         | <https://pkg.go.dev/github.com/fxamacker/cbor/v2>                   |
| `datastore`    | <https://pkg.go.dev/cloud.google.com/go/datastore>                      |
| `default`      | The type of "default" must match the type of the field.                 |
| `json`         | <https://pkg.go.dev/encoding/json>                                      |
| `mapstructure` | <https://pkg.go.dev/github.com/mitchellh/mapstructure>                  |
| `properties`   | <https://pkg.go.dev/github.com/magiconair/properties#Properties.Decode> |
| `protobuf`     | <https://github.com/golang/protobuf>                                    |
| `required`     | Should be only "true" or "false".                                       |
| `spanner`      | <https://pkg.go.dev/cloud.google.com/go/spanner>                        |
| `toml`         | <https://pkg.go.dev/github.com/pelletier/go-toml/v2>                    |
| `url`          | <https://github.com/google/go-querystring>                              |
| `validate`     | <https://github.com/go-playground/validator>                            |
| `xml`          | <https://pkg.go.dev/encoding/xml>                                       |
| `yaml`         | <https://pkg.go.dev/gopkg.in/yaml.v2>                                   |

_Configuration_: (optional) The list of struct tags that can be accepted by the rule additionally to the supported tags.

Configuration example:

To accept the `inline` option in JSON tags (and `outline` and `gnu` in BSON tags) you must provide the following configuration

```toml
[rule.struct-tag]
arguments = ["json,inline", "bson,outline,gnu"]
```

To prevent a tag from being checked, simply add a `!` before its name.
For example, to instruct the rule not to check `validate` tags
(and accept `outline` and `gnu` in BSON tags) you can provide the following configuration

```toml
[rule.struct-tag]
arguments = ["!validate", "bson,outline,gnu"]
```

