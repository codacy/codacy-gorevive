## datarace

_Description_: This rule spots potential dataraces caused by goroutines capturing (by-reference) particular identifiers of the function from
which goroutines are created.
The rule is able to spot two of such cases: go-routines capturing named return values, and capturing `for-range` values.

_Configuration_: N/A

