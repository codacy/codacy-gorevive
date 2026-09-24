## receiver-naming

_Go version_: 1.0.

**_Ported from golint_**

_Description_: By convention, receiver names in a method should reflect their identity.
For example, if the receiver is of type `Parts`, `p` is an adequate name for it.
Contrary to other languages, it is not idiomatic to name receivers as `this` or `self`.
All methods of a type should also use the same receiver name.

