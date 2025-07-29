## enforce-switch-style

_Description_: This rule enforces consistent usage of `default` on `switch` statements.
It can check for `default` case clause occurrence and/or position in the list of case clauses.

_Configuration_: ([]string) Specifies what to enforced: occurrence and/or position. The, non-mutually exclusive, options are:

- "allowNoDefault": allows `switch` without `default` case clause.
- "allowDefaultNotLast": allows `default` case clause to be not the last clause of the `switch`.

Examples:

To enforce that all `switch` statements have a `default` clause as its the last case clause:

```toml
[rule.enforce-switch-style]
```

To enforce that all `switch` statements have a `default` clause but its position is unimportant:

```toml
[rule.enforce-switch-style]
arguments = ["allowDefaultNotLast"]
```

To enforce that in all `switch` statements with a `default` clause, the `default` is the last case clause:

```toml
[rule.enforce-switch-style]
arguments = ["allowNoDefault"]
```

Notice that a configuration including both options will effectively deactivate the whole rule.

