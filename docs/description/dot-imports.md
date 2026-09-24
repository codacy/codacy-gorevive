## dot-imports

_Go version_: 1.0.

**_Ported from golint_**

_Description_: Importing with `.` makes the programs much harder to understand because it is unclear whether names belong to the current package or
to an imported package.

More [information here](https://go.dev/wiki/CodeReviewComments#import-dot).

_Configuration_:

- `allowed-packages`: (list of strings) list of allowed dot import packages

Configuration example:

```toml
[rule.dot-imports]
arguments = [
  { allowed-packages = [
    "github.com/onsi/ginkgo/v2",
    "github.com/onsi/gomega",
  ] },
]
```

