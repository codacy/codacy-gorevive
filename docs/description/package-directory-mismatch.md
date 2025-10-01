## package-directory-mismatch

_Description_: It is considered a good practice to name a package after the directory containing it.
This rule warns when the package name declared in the file does not match the name of the directory containing the file.

The following cases are excluded from this check:

- Package `main` (executable packages)
- Files in `testdata` directories (at any level) - by default
- Files directly in `internal` directories (but files in subdirectories of `internal` are checked)

For test files (files with `_test` suffix), package name additionally checked if it matches directory name with  `_test` suffix appended.

The rule normalizes both directory and package names before comparison by removing hyphens (`-`),
underscores (`_`), and dots (`.`). This allows package `foo_barbuz` be equal with directory `foo-bar.buz`.

For files in version directories (`v1`, `v2`, etc.), package name is checked if it matches either the version directory or its parent directory.

_Configuration_: Named arguments for directory exclusions.

Examples:

Default behavior excludes paths containing `testdata`

```toml
[rule.package-directory-mismatch]
```

Ignore specific directories with `ignore-directories`

```toml
[rule.package-directory-mismatch]
arguments = [{ ignore-directories = ["testcases", "testinfo"] }]
```

Include all directories (`testdata` also)

```toml
[rule.package-directory-mismatch]
arguments = [{ ignoreDirectories = [] }]
```

