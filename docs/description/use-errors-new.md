## use-errors-new

_Go version_: < 1.26.

_Description_: This rule identifies calls to `fmt.Errorf` that can be safely replaced by, the more efficient, `errors.New`.
This applies when the format string has no formatting verbs (no additional arguments are passed).

