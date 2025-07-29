## file-length-limit

_Description_: This rule enforces a maximum number of lines per file, in order to aid in maintainability and reduce complexity.

_Configuration_:

- `max`: (int) a maximum number of lines in a file. Must be non-negative integers. 0 means the rule is disabled (default `0`);
- `skipComments` (`skipcomments`, `skip-comments`): (bool) if true ignore and do not count lines containing just comments (default `false`);
- `skipBlankLines` (`skipblanklines`, `skip-blank-lines`): (bool) if true ignore and do not count lines made up purely of whitespace (default `false`).

Examples:

```toml
[rule.file-length-limit]
arguments = [{ max = 100, skipComments = true, skipBlankLines = true }]
```

```toml
[rule.file-length-limit]
arguments = [{ max = 100, skip-comments = true, skip-blank-lines = true }]
```

