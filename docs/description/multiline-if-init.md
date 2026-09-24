## multiline-if-init

_Go version_: 1.0.

_Description_: Flags `if` statements whose init clause spans multiple lines.
The if-init idiom exists for tight one-liners.
When the init wraps across lines, the reader has to visually parse a struct literal or
call chain to find where the initialization ends and the condition begins.
Extract the initialization to a separate statement instead.

