## optimize-operands-order

_Description_: Conditional expressions can be written to take advantage of short circuit evaluation and speed up its average evaluation time
by forcing the evaluation of less time-consuming terms before more costly ones.
This rule spots logical expressions where the order of evaluation of terms seems non optimal.
Please notice that confidence of this rule is low and is up to the user to decide if the suggested rewrite of the expression
keeps the semantics of the original one.

_Configuration_: N/A

Example:

```golang
if isGenerated(content) && !config.IgnoreGeneratedHeader {
```

Swap left and right side :

```golang
if !config.IgnoreGeneratedHeader && isGenerated(content) {
```

