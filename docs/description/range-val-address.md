## range-val-address

_Go version_: < 1.22.

**_Typed_**

_Description_: Range variables in a loop are reused at each iteration.
This rule warns when assigning the address of the variable, passing the address to append() or using it in a map.

_Configuration_: N/A

_Note_: This rule is irrelevant for Go 1.22+.

