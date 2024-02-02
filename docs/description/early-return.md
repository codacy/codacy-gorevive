## early-return

_Description_: In Go it is idiomatic to minimize nesting statements, a typical example is to avoid if-then-else constructions. This rule spots constructions like
```go
if cond {
  // do something
} else {
  // do other thing
  return ...
}
```
where the `if` condition may be inverted in order to reduce nesting:
```go
if !cond {
  // do other thing
  return ...
}

// do something
```

_Configuration_: ([]string) rule flags. Available flags are:

* _preserveScope_: do not suggest refactorings that would increase variable scope

Example:

```toml
[rule.early-return]
  arguments = ["preserveScope"]
```

