## string-format

_Description_: This rule allows you to configure a list of regular expressions that string literals in certain function calls are checked against.
This is geared towards user facing applications where string literals are often used for messages that will be presented to users, so it may be desirable to enforce consistent formatting.

_Configuration_: Each argument is a slice containing 2-3 strings: a scope, a regex, and an optional error message.

1. The first string defines a **scope**. This controls which string literals the regex will apply to, and is defined as a function argument. It must contain at least a function name (`core.WriteError`). Scopes may optionally contain a number specifying which argument in the function to check (`core.WriteError[1]`), as well as a struct field (`core.WriteError[1].Message`, only works for top level fields). Function arguments are counted starting at 0, so `[0]` would refer to the first argument, `[1]` would refer to the second, etc. If no argument number is provided, the first argument will be used (same as `[0]`).

2. The second string is a **regular expression** (beginning and ending with a `/` character), which will be used to check the string literals in the scope. The default semantics is "_strings matching the regular expression are OK_". If you need to inverse the semantics you can add a `!` just before the first `/`. Examples:

    * with `"/^[A-Z].*$/"` the rule will **accept** strings starting with capital letters
    * with `"!/^[A-Z].*$/"` the rule will a **fail** on strings starting with capital letters

3. The third string (optional) is a **message** containing the purpose for the regex, which will be used in lint errors.

Example:

```toml
[rule.string-format]
  arguments = [
    ["core.WriteError[1].Message", "/^([^A-Z]|$)/", "must not start with a capital letter"],
    ["fmt.Errorf[0]", "/(^|[^\\.!?])$/", "must not end in punctuation"],
    ["panic", "/^[^\\n]*$/", "must not contain line breaks"]]
```

