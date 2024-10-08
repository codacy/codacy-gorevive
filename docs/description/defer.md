## defer

_Description_: This rule warns on some common mistakes when using `defer` statement. It currently alerts on the following situations:

| name              | description                                                                                                                                                                                     |
|-------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| call-chain        | even if deferring call-chains of the form `foo()()` is valid, it does not helps code understanding (only the last call is deferred)                                                             |
| loop              | deferring inside loops can be misleading (deferred functions are not executed at the end of the loop iteration but of the current function) and it could lead to exhausting the execution stack |
| method-call       | deferring a call to a method can lead to subtle bugs if the method does not have a pointer receiver                                                                                             |
| recover           | calling `recover` outside a deferred function has no effect                                                                                                                                     |
| immediate-recover | calling `recover` at the time a defer is registered, rather than as part of the deferred callback.  e.g. `defer recover()` or equivalent.                                                       |
| return            | returning values form a deferred function has no effect                                                                                                                                         |

These gotchas are described [here](https://blog.learngoprogramming.com/gotchas-of-defer-in-go-1-8d070894cb01)

_Configuration_: by default all warnings are enabled but it is possible selectively enable them through configuration. For example to enable only `call-chain` and `loop`:

```toml
[rule.defer]
  arguments = [["call-chain","loop"]]
```

