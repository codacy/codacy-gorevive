## forbidden-call-in-wg-go

_Description_: Since Go 1.25, it is possible to create goroutines with the method `waitgroup.Go`.
The `Go` method calls a function in a new goroutine and adds (`Add`) that task to the WaitGroup.
When the function returns, the task is removed (`Done`) from the WaitGroup.

This rule ensures that functions don't panic as is specified
in the [documentation of `WaitGroup.Go`](https://pkg.go.dev/sync#WaitGroup.Go).

The rule also warns against a common mistake when refactoring legacy code:
accidentally leaving behind a call to `WaitGroup.Done`, which can cause subtle bugs or panics.

