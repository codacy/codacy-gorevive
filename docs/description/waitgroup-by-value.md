## waitgroup-by-value

_Description_: Function parameters that are passed by value, are in fact a copy of the original argument. Passing a copy of a `sync.WaitGroup` is usually not what the developer wants to do.
This rule warns when a `sync.WaitGroup` expected as a by-value parameter in a function or method.

_Configuration_: N/A
