## increment-decrement

_Go version_: 1.0.

**_Ported from golint_**

_Description_: By convention, for better readability, incrementing an integer variable by 1 is recommended to be done using the `++` operator.
This rule spots expressions like `i += 1` and `i -= 1` and proposes to change them into `i++` and `i--`.

