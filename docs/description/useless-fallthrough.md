## useless-fallthrough

_Description_: This rule warns on useless `fallthrough` statements in case clauses of switch statements.
A `fallthrough` is considered _useless_ if it's the single statement of a case clause block.

Go allows `switch` statements with clauses that group multiple cases.
Thus, for example:

```go
switch category {
  case "Lu":
    fallthrough
  case "Ll":
    fallthrough    
  case "Lt":
    fallthrough
  case "Lm": 
    fallthrough
  case "Lo":
    return true
  default:
    return false
}
```

can be written as

```go
switch category {
  case "Lu", "Ll", "Lt", "Lm", "Lo":
      return true
  default:
    return false
}
```

_Configuration_: N/A

