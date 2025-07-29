## early-return

_Description_: In Go it is idiomatic to minimize nesting statements, a typical example is to avoid if-then-else constructions.
This rule spots constructions like

```golang
if cond {
	// do something
} else {
	// do other thing
	return ...
}
```

where the `if` condition may be inverted in order to reduce nesting:

```golang
if !cond {
	// do other thing
	return ...
}

// do something
```

_Configuration_: ([]string) rule flags. Available flags are:

- `preserveScope` (`preservescope`, `preserve-scope`): do not suggest refactorings that would increase variable scope
- `allowJump` (`allowjump`, `allow-jump`): suggest a new jump (`return`, `continue` or `break` statement) if it could unnest multiple statements.
By default, only relocation of _existing_ jumps (i.e. from the `else` clause) are suggested.

Examples:

```toml
[rule.early-return]
arguments = ["preserveScope", "allowJump"]
```

```toml
[rule.early-return]
arguments = ["preserve-scope", "allow-jump"]
```

