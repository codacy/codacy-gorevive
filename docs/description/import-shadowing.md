## import-shadowing

_Description_: In Go it is possible to declare identifiers (packages, structs,
interfaces, parameters, receivers, variables, constants...) that conflict with the
name of an imported package. This rule spots identifiers that shadow an import.

The rule ignores versioned import paths such as `k8s.io/api/core/v1` when `v1` is the package name,
which allows identifiers like `v1`. This is a deliberate trade-off to keep the rule simple.

_Configuration_: N/A

