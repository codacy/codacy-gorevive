## context-as-argument

_Description_: By [convention](https://go.dev/wiki/CodeReviewComments#contexts), `context.Context` should be the first parameter of a function. This rule spots function declarations that do not follow the convention.

_Configuration_:

* `allowTypesBefore` : (string) comma-separated list of types that may be before 'context.Context'

Example:

```toml
[rule.context-as-argument]
  arguments = [{allowTypesBefore = "*testing.T,*github.com/user/repo/testing.Harness"}]
```

