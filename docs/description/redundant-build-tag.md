## redundant-build-tag

_Description_: This rule warns about redundant build tag comments `// +build` when `//go:build` is present.
`gofmt` in Go 1.17+ automatically adds the `//go:build` constraint, making the `// +build` comment unnecessary.

_Configuration_: N/A

