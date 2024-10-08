## dot-imports

_Description_: Importing with `.` makes the programs much harder to understand because it is unclear whether names belong to the current package or to an imported package.

More information [here](https://go.dev/wiki/CodeReviewComments#import-dot)

_Configuration_:

* `allowedPackages`: (list of strings) comma-separated list of allowed dot import packages

Example:

```toml
[rule.dot-imports]
  arguments = [{ allowedPackages = ["github.com/onsi/ginkgo/v2","github.com/onsi/gomega"] }]
```

