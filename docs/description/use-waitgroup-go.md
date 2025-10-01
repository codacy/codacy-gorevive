## use-waitgroup-go

_Description_: Since Go 1.25 the `sync` package proposes the [`WaitGroup.Go`](https://pkg.go.dev/sync#WaitGroup.Go) method.
This method is a shorter and safer replacement for the idiom `wg.Add ... go { ... wg.Done ... }`.
The rule proposes to replace these legacy idioms with calls to the new method.

_Limitations_: The rule doesn't rely on type information but on variable names to identify waitgroups.
This means the rule search for `wg` (the defacto standard name for wait groups);
if the waitgroup variable is named differently than `wg` the rule will skip it.

_Configuration_: N/A

