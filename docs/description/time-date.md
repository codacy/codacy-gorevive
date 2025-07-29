## time-date

_Description_: Reports bad usage of `time.Date`.

_Configuration_: N/A

_Examples_:

- Invalid dates reporting:

  - 0 for the month or day argument
  - out of bounds argument for the month (12), day (31), hour (23), minute (59), or seconds (59)
  - an invalid date: 31st of June, 29th of February in 2023, ...

- Non-decimal integers are used as arguments

  This includes:

  - leading zero notation like using 00 for hours, minutes, and seconds.
  - octal notation 0o1, 0o0 that are often caused by using gofumpt on leading zero notation.
  - padding zeros such as 00123456 that are source of bugs.
  - ... and some other use cases.

<details>
<summary>More information about what is detected and reported</summary>

```go
import "time"

var (
	// here we can imagine zeroes were used for padding purpose
	a = time.Date(2023, 1, 2, 3, 4, 0, 00000000, time.UTC) // 00000000 is octal and equals 0 in decimal
	b = time.Date(2023, 1, 2, 3, 4, 0, 00000006, time.UTC) // 00000006 is octal and equals 6 in decimal
	c = time.Date(2023, 1, 2, 3, 4, 0, 00123456, time.UTC) // 00123456 is octal and equals 42798 in decimal
)
```

But here, what was expected 123456 or 42798 ? It's a source of bugs.

So the rule will report this

> octal notation with padding zeroes found: choose between 123456 and 42798 (decimal value of 123456 octal value)

This rule also reports strange notations used with time.Date.

Example:

```go
import "time"

var _ = time.Date(
	0x7e7,    // hexadecimal notation: use 2023 instead of 0x7e7/
	0b1,      // binary notation: use 1 instead of 0b1/
	0x_2,     // hexadecimal notation: use 2 instead of 0x_2/
	1_3,      // alternative notation: use 13 instead of 1_3/
	1e1,      // exponential notation: use 10 instead of 1e1/
	0.,       // float literal: use 0 instead of 0./
	0x1.Fp+6, // float literal: use 124 instead of 0x1.Fp+6/
	time.UTC)
```

All these are considered to be an uncommon usage of time.Date, are reported with a 0.8 confidence.

Note: even if 00, 01, 02, 03, 04, 05, 06, 07 are octal numbers, they can be considered as valid, and reported with 0.5 confidence.

```go
import "time"

var _ = time.Date(2023, 01, 02, 03, 04, 00, 0, time.UTC)
```

</details>

