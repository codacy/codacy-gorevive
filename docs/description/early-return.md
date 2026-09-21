## early-return

_Go version_: 1.0.

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

