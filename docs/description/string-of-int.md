## string-of-int
_Description_:  explicit type conversion `string(i)` where `i` has an integer type other than `rune` might behave not as expected by the developer (e.g. `string(42)` is not `"42"`). This rule spot that kind of suspicious conversions. 

_Configuration_: N/A

