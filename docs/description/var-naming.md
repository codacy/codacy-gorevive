## var-naming

_Go version_: 1.0.

**_Ported from golint_**

_Description_: This rule warns when [initialism](https://go.dev/wiki/CodeReviewComments#initialisms), [variable](https://go.dev/wiki/CodeReviewComments#variable-names)
naming conventions are not followed.
It ignores functions starting with `Example`, `Test`, `Benchmark`, and `Fuzz` in test files, preserving `golint` original behavior.

_Configuration_: This rule accepts two slices of strings and one optional slice containing a single map with named parameters.
(This is because TOML does not support "slice of any," and we maintain backward compatibility with the previous configuration version).
The first slice is an allowlist, and the second one is a blocklist of initialisms.
You can add a boolean parameter `skip-initialism-name-checks` to control how names
of functions, variables, consts, and structs handle known initialisms (e.g., JSON, HTTP, etc.) when written in `camelCase`.
When `skip-initialism-name-checks` is set to true, the rule allows names like `readJson`, `HttpMethod` etc.
In the map, you can add a boolean `upper-case-const` parameter to allow `UPPER_CASE` for `const`.

By default, the rule behaves exactly as the alternative in `golint` for non-package identifiers;
`golint`-equivalent package-name warnings now require enabling the [`package-naming`](#package-naming) rule.
The legacy package-related options `skip-package-name-checks`, `extra-bad-package-names`, and `skip-package-name-collision-with-go-std` are deprecated
and are now treated as no-ops by `var-naming` (they are ignored, apart from an optional warning when logging is enabled).
Package-name checks should be configured via the [`package-naming`](#package-naming) rule instead,
and these options should be removed from `var-naming` configurations to avoid confusion.

Configuration examples:

```toml
[rule.var-naming]
arguments = [[], [], [{ skip-initialism-name-checks = true }]]
```

```toml
[rule.var-naming]
arguments = [["ID"], ["VM"], [{ upper-case-const = true }]]
```

