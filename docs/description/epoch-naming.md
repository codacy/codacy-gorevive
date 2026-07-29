## epoch-naming

_Description_: Variables initialized with epoch time methods (`time.Now().Unix()`, `time.Now().UnixMilli()`,
`time.Now().UnixMicro()`, `time.Now().UnixNano()`) should have names that clearly indicate their time unit to
prevent confusion and potential bugs when working with different time scales.

This rule enforces that variable names contain appropriate suffixes based on the method used:

- `Unix()`: variable name should end with "Sec", "Second" or "Seconds"
- `UnixMilli()`: variable name should end with "Milli" or "Ms"
- `UnixMicro()`: variable name should end with "Micro", "Microsecond", "Microseconds" or "Us"
- `UnixNano()`: variable name should end with "Nano" or "Ns"

The rule checks variable declarations, short variable declarations (`:=`), and regular assignments (`=`).
The suffix matching is case-insensitive and must appear at the end of the variable name.

