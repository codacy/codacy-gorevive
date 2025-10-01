## modifies-parameter

_Description_: A function that modifies its parameters can be hard to understand.
It can also be misleading if the arguments are passed by value by the caller.
This rule warns when a function modifies one or more of its parameters or when
parameters are passed to functions that modify them (e.g. `slices.Delete`).

_Configuration_: N/A

