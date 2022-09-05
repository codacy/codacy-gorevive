## datarace

_Description_: This rule spots potential dataraces caused by go-routines capturing (by-reference) particular identifiers of the function from which go-routines are created. The rule is able to spot two of such cases: go-routines capturing named return values, and capturing `for-range` values.

_Configuration_: N/A

