## file-header

_Description_: This rule helps to enforce a common header for all source files in a project by spotting those files that do not have the specified header.

_Configuration_: (string) the header to look for in source files.

Example:

```toml
[rule.file-header]
  arguments =["This is the text that must appear at the top of source files."]
```

