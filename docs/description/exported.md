## exported

_Description_: Exported function and methods should have comments. This warns on undocumented exported functions and methods.

More information [here](https://github.com/golang/go/wiki/CodeReviewComments#doc-comments)

_Configuration_: ([]string) rule flags.
Please notice that without configuration, the default behavior of the rule is that of its `golint` counterpart.
Available flags are:

* _checkPrivateReceivers_ enables checking public methods of private types
* _disableStutteringCheck_ disables checking for method names that stutter with the package name (i.e. avoid failure messages of the form _type name will be used as x.XY by other packages, and that stutters; consider calling this Y_)
* _sayRepetitiveInsteadOfStutters_ replaces the use of the term _stutters_ by _repetitive_ in failure messages

Example:

```toml
[rule.exported]
  arguments =["checkPrivateReceivers","disableStutteringCheck"]
```

