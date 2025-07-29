## var-naming

_Description_: This rule warns when [initialism](https://go.dev/wiki/CodeReviewComments#initialisms), [variable](https://go.dev/wiki/CodeReviewComments#variable-names)
or [package](https://go.dev/wiki/CodeReviewComments#package-names) naming conventions are not followed.
It ignores functions starting with `Example`, `Test`, `Benchmark`, and `Fuzz` in test files, preserving `golint` original behavior.

_Configuration_: This rule accepts two slices of strings and one optional slice containing a single map with named parameters.
(This is because TOML does not support "slice of any," and we maintain backward compatibility with the previous configuration version).
The first slice is an allowlist, and the second one is a blocklist of initialisms.
You can add a boolean parameter `skipInitialismNameChecks` (`skipinitialismnamechecks` or `skip-initialism-name-checks`) to control how names
of functions, variables, consts, and structs handle known initialisms (e.g., JSON, HTTP, etc.) when written in `camelCase`.
When `skipInitialismNameChecks` is set to true, the rule allows names like `readJson`, `HttpMethod` etc.
In the map, you can add a boolean `upperCaseConst` (`uppercaseconst`, `upper-case-const`) parameter to allow `UPPER_CASE` for `const`.
You can also add a boolean `skipPackageNameChecks` (`skippackagenamechecks`, `skip-package-name-checks`) to skip package name checks.
When `skipPackageNameChecks` is false (the default), you can configure `extraBadPackageNames` (`extrabadpackagenames`, `extra-bad-package-names`)
to forbid using the values from the list as package names additionally to the standard meaningless ones:
"common", "interfaces", "misc", "types", "util", "utils".

By default, the rule behaves exactly as the alternative in `golint` but optionally, you can relax it (see [golint/lint/issues/89](https://github.com/golang/lint/issues/89)).

Examples:

```toml
[rule.var-naming]
arguments = [[], [], [{ skipInitialismNameChecks = true }]]
```

```toml
[rule.var-naming]
arguments = [["ID"], ["VM"], [{ upperCaseConst = true }]]
```

```toml
[rule.var-naming]
arguments = [[], [], [{ skipPackageNameChecks = true }]]
```

```toml
[rule.var-naming]
arguments = [[], [], [{ extraBadPackageNames = ["helpers", "models"] }]]
```

```toml
[rule.var-naming]
arguments = [[], [], [{ skip-initialism-name-checks = true }]]
```

```toml
[rule.var-naming]
arguments = [["ID"], ["VM"], [{ upper-case-const = true }]]
```

```toml
[rule.var-naming]
arguments = [[], [], [{ skip-package-name-checks = true }]]
```

```toml
[rule.var-naming]
arguments = [[], [], [{ extra-bad-package-names = ["helpers", "models"] }]]
```

