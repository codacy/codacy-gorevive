## marshal-receiver

_Go version_: 1.0.

_Description_: Checks receiver type consistency for common marshal/unmarshal methods.
The rule inspects only methods whose names exactly match: `MarshalJSON`, `MarshalText`, `MarshalYAML`, `UnmarshalJSON`, `UnmarshalText`, and `UnmarshalYAML`.
For these methods, it enforces receiver kind only:

- `Marshal*` methods should use a value receiver, and are reported when declared with a pointer receiver.
- `Unmarshal*` methods should use a pointer receiver, and are reported when declared with a value receiver.

This is a name-based, syntactic check.
It does not validate method signatures (parameters or return values) and does not verify whether a method satisfies a specific marshaling interface.

