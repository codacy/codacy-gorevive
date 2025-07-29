## modifies-value-receiver

_Description_: A method that modifies its receiver value can have undesired behavior.
The modification can be also the root of a bug because the actual value receiver could be a copy of that used at the calling site.
This rule warns when a method modifies its receiver.

_Configuration_: N/A

