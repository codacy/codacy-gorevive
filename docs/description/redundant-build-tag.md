## redundant-build-tag

_Go version_: >= 1.17; behavior changes in 1.21.

_Description_: This rule warns about redundant [build tag comments](https://pkg.go.dev/cmd/go@go1.17.0#hdr-Build_constraints).
It detects unnecessary `// +build` comments when `//go:build` is present.
`gofmt` in Go 1.17+ automatically adds the `//go:build` constraint, making the `// +build` comment unnecessary.
Also, the rule spots redundant build tags `//go:build go1.X` when the package's Go language version is greater than or equal to `go1.X`.

