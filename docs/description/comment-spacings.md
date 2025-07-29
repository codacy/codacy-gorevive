## comment-spacings

_Description_: Spots comments of the form:

```go
//This is a malformed comment: no space between // and the start of the sentence
```

_Configuration_: ([]string) list of exceptions. For example, to accept comments of the form

```go
//mypragma: activate something
//+optional
```

You need to add both `"mypragma:"` and `"+optional"` in the configuration

Example:

```toml
[rule.comment-spacings]
arguments = ["mypragma:", "+optional"]
```

