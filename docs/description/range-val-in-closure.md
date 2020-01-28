## range-val-in-closure

_Description_: Range variables in a loop are reused at each iteration; therefore a goroutine created in a loop will point to the range variable with from the upper scope. This way, the goroutine could use the variable with an undesired value.
This rule warns when a range value (or index) is used inside a closure

_Configuration_: N/A
