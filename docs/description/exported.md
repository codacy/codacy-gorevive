## exported

_Description_: Exported function and methods should have comments. This warns on undocumented exported functions and methods.

More [information here](https://go.dev/wiki/CodeReviewComments#doc-comments).

_Configuration_: ([]string) rule flags.
Please notice that without configuration, the default behavior of the rule is that of its `golint` counterpart.
Available flags are:

- `checkPrivateReceivers` (`checkprivatereceivers`, `check-private-receivers`) enables checking public methods of private types
- `disableStutteringCheck` (`disablestutteringcheck`, `disable-stuttering-check`) disables checking for method names that stutter with the package name
(i.e. avoid failure messages of the form _type name will be used as x.XY by other packages, and that stutters; consider calling this Y_)
- `sayRepetitiveInsteadOfStutters` (`sayrepetitiveinsteadofstutters`, `say-repetitive-instead-of-stutters`) replaces the use of the term _stutters_
by _repetitive_ in failure messages
- `checkPublicInterface` (`checkpublicinterface`, `check-public-interface`) enabled checking public method definitions in public interface types
- `disableChecksOnConstants` (`disablechecksonconstants`, `disable-checks-on-constants`) disable all checks on constant declarations
- `disableChecksOnFunctions` (`disablechecksonfunctions`, `disable-checks-on-functions`) disable all checks on function declarations
- `disableChecksOnMethods` (`disablechecksonmethods`, `disable-checks-on-methods`) disable all checks on method declarations
- `disableChecksOnTypes` (`disablechecksontypes`, `disable-checks-on-types`) disable all checks on type declarations
- `disableChecksOnVariables` (`disablechecksonvariables`, `disable-checks-on-variables`) disable all checks on variable declarations

Examples:

```toml
[rule.exported]
arguments = [
  "checkPrivateReceivers",
  "disableStutteringCheck",
  "checkPublicInterface",
  "disableChecksOnFunctions",
]
```

```toml
[rule.exported]
arguments = [
  "check-private-receivers",
  "disable-stuttering-check",
  "check-public-interface",
  "disable-checks-on-functions",
]
```

