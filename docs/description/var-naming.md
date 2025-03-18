## var-naming

_Description_: This rule warns when [initialism](https://go.dev/wiki/CodeReviewComments#initialisms), [variable](https://go.dev/wiki/CodeReviewComments#variable-names) or [package](https://go.dev/wiki/CodeReviewComments#package-names) naming conventions are not followed. It ignores functions starting with `Example`, `Test`, `Benchmark`, and `Fuzz` in test files, preserving `golint` original behavior.

_Configuration_: This rule accepts two slices of strings and one optional slice with single map with named parameters.
(it's due to TOML hasn't "slice of any" and we keep backward compatibility with previous config version)
First slice is an allowlist and second one is a blocklist of initialisms.
In map, you can add "upperCaseConst=true" parameter to allow `UPPER_CASE` for `const`
By default, the rule behaves exactly as the alternative in `golint` but optionally, you can relax it (see [golint/lint/issues/89](https://github.com/golang/lint/issues/89))

Example:

```toml
[rule.var-naming]
  arguments = [["ID"], ["VM"], [{upperCaseConst=true}]]
```

You can also add "skipPackageNameChecks=true" to skip package name checks.

Example:


```toml
[rule.var-naming]
  arguments = [[], [], [{skipPackageNameChecks=true}]]
```

