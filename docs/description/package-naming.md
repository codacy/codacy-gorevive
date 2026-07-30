## package-naming

_Description_: This rule checks that package names follow [Go conventions](https://go.dev/blog/package-names) and best practices.
It helps prevent using bad package names and enforces consistent naming patterns.
This rule arose from package naming checks in `var-naming`.

By default, it checks for:

- Package name conventions (no underscores except for test packages, no MixedCaps).
- Bad package names from the official Go blog (e.g., `common`, `util`, `utils`, `misc`, `interfaces`, `types`).
- Package names that conflict with common Go standard library packages (e.g., `http`, `json`, `fmt`).

_Configuration_: (optional) single map of options (`map[string]any`), provided as the single configuration argument in the rule's arguments array.

- `skipConventionNameCheck` (`skipconventionnamecheck`, `skip-convention-name-check`): (bool)
If `true`, skip checks for package name conventions (underscores, MixedCaps, etc.). Default: `false`.
This option is mutually exclusive with `conventionNameCheckRegex`; setting both results in a configuration error.
- `conventionNameCheckRegex` (`conventionnamecheckregex`, `convention-name-check-regex`): (string)
Custom regex pattern to validate package names. If set, package names must match this pattern.
The value must be a non-empty string, and this option is mutually exclusive with `skipConventionNameCheck`; setting both results in a configuration error.
- `skipTopLevelCheck` (`skiptoplevelcheck`, `skip-top-level-check`): (bool)
If `true`, skip checks for top-level package names (e.g., `pkg`). Default: `false`.
- `skipDefaultBadNameCheck` (`skipdefaultbadnamecheck`, `skip-default-bad-name-check`): (bool)
If `true`, skip checks for default bad package names (e.g., `common`, `utils`). Default: `false`.
- `checkExtraBadName` (`checkextrabadname`, `check-extra-bad-name`): (bool)
If `true`, enable checks for extra bad package names (e.g., `helpers`, `models`, `shared`, `utilities`). Default: `false`.
- `userDefinedBadNames` (`userdefinedbadnames`, `user-defined-bad-names`): ([]string) List of user-defined bad package names to check for.
- `skipCollisionWithCommonStd` (`skipcollisionwithcommonstd`, `skip-collision-with-common-std`): (bool)
If `true`, skip checks for collisions with the most common Go standard library packages. Default: `false`.
- `checkCollisionWithAllStd` (`checkcollisionwithallstd`, `check-collision-with-all-std`): (bool)
If `true`, enable checks for collisions with all packages from Go standard library. Default: `false`.
This option is mutually exclusive with `skipCollisionWithCommonStd`; setting both results in a configuration error.

Configuration examples:

Default settings (check for name conventions, top level packages, common bad names, and collision with common Go standard library packages):

```toml
[rule.package-naming]
```

Custom naming convention with regex:

```toml
[rule.package-naming]
arguments = [{ convention-name-check-regex = "^[a-z]+$" }]
```

Skip convention checks, but check for bad names:

```toml
[rule.package-naming]
arguments = [{ skip-convention-name-check = true }]
```

Enable collision checks with the most common standard library packages:

```toml
[rule.package-naming]
arguments = [{ skip-collision-with-common-std = false }]
```

Strict mode with user-defined bad names:

```toml
[rule.package-naming]
arguments = [{ user-defined-bad-names = ["foo", "bar"] }]
```

Enable collision checks with all standard library [packages](https://pkg.go.dev/std) excluding `internal` and `vendor`:

```toml
[rule.package-naming]
arguments = [{ check-collision-with-all-std = true }]
```

