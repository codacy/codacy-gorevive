## unnecessary-if

_Go version_: 1.0.

_Description_: Detects unnecessary `if-else` statements that return or assign a boolean value
based on a condition and suggests a simplified, direct return or assignment.
The `if-else` block is redundant because the condition itself is already a boolean expression.
The simplified version is immediately clearer, more idiomatic, and reduces cognitive load for the reader.

