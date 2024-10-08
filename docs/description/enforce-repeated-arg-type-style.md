## enforce-repeated-arg-type-style

_Description_: This rule is designed to maintain consistency in the declaration
of repeated argument and return value types in Go functions. It supports three styles:
'any', 'short', and 'full'. The 'any' style is lenient and allows any form of type
declaration. The 'short' style encourages omitting repeated types for conciseness,
whereas the 'full' style mandates explicitly stating the type for each argument
and return value, even if they are repeated, promoting clarity.

_Configuration (1)_: (string) as a single string, it configures both argument
and return value styles. Accepts 'any', 'short', or 'full' (default: 'any').

_Configuration (2)_: (map[string]any) as a map, allows separate configuration
for function arguments and return values. Valid keys are "funcArgStyle" and
"funcRetValStyle", each accepting 'any', 'short', or 'full'. If a key is not
specified, the default value of 'any' is used.

_Note_: The rule applies checks based on the specified styles. For 'full' style,
it flags instances where types are omitted in repeated arguments or return values.
For 'short' style, it highlights opportunities to omit repeated types for brevity.
Incorrect or unknown configuration values will result in an error.

Example (1):

```toml
[rule.enforce-repeated-arg-type-style]
  arguments = ["short"]
```

Example (2):

```toml
[rule.enforce-repeated-arg-type-style]
  arguments = [{ funcArgStyle = "full", funcRetValStyle = "short" }]
```

