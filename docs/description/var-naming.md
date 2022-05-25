## var-naming

_Description_: This rule warns when [initialism](https://github.com/golang/go/wiki/CodeReviewComments#initialisms), [variable](https://github.com/golang/go/wiki/CodeReviewComments#variable-names) or [package](https://github.com/golang/go/wiki/CodeReviewComments#package-names) naming conventions are not followed.

_Configuration_: This rule accepts two slices of strings, a whitelist and a blacklist of initialisms. By default, the rule behaves exactly as the alternative in `golint` but optionally, you can relax it (see [golint/lint/issues/89](https://github.com/golang/lint/issues/89))

Example:

```toml
[rule.var-naming]
  arguments = [["ID"], ["VM"]]
```

