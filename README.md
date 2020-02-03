# Codacy Revive

Codacy tool to add support for [revive](https://github.com/mgechev/revive).

# Usage

You can create the docker by doing:

```
make docker
```

The docker is ran with the following command:

```
docker run -it -v $srcDir:/src <DOCKER_NAME>:<DOCKER_VERSION>
```

To run the tool using a custom configuration file, run docker with the following command:

```
docker run -it -v $srcDir:/src -v $codacyrcConfig:/.codacyrc <DOCKER_NAME>:<DOCKER_VERSION>
```

## Tool Developer Documentation

[Tool Developer Guide](https://support.codacy.com/hc/en-us/articles/207994725-Tool-Developer-Guide)


[Tool Developer Guide - Using Scala](https://support.codacy.com/hc/en-us/articles/207280379-Tool-Developer-Guide-Using-Scala)


## Structure

- To run the tool we provide the configuration file, `/.codacyrc`, with the language to run and optional parameters your tool might need.
- The source code to be analysed will be located in `/src`, meaning that when provided in the configuration, the file paths are relative to `/src`.

#### Configuration file (.codacyrc)

This file has:

 - files: Files to be analysed (their path is relative to `/src`)
 - tools: Array of tools
 - name: Unique identifier of the tool. This will be provided by the tool in patterns.json file.
 - patterns: Array of patterns that must be checked
     - patternId: Unique identifier of the pattern
     - parameters: Parameters of the pattern
     - name: Unique identifier of the parameter
     - value: Value to be used as parameter value

```
{
  "files" : ["foo/bar/baz.go", "foo2/bar/baz.go"],
  "tools":[
    {
      "name":"codacy-gorevive",
      "patterns":[
        {
          "patternId":"exported"
        }
      ]
    }
  ]
}
```

#### Tool documentation

At Codacy we strive to provide the best value to our users and, to accomplish that, we document our patterns so that the user can better understand the problem and fix it.

The documentation for the tool must always be updated before submitting the docker.

Your files for this section should be placed in `/docs/description/`.

In order to provide more details you can create:

- A single `/docs/description/description.json`
- A `/docs/description/<PATTERN-ID>.md` for each pattern


In the description.json you define the title for the pattern, brief description, time to fix (in minutes), and also a description of the parameters in the following format:

```
[
  {
    "patternId": "add-constant",
    "title": "add-constant",
    "description": "Suggests using constant for [magic numbers](https://en.wikipedia.org/wiki/Magic_number_(programming)#Unnamed_numerical_constants) and string literals.",
    "parameters": [
      {
        "name": "allowFloats",
        "default": "",
        "description": "(string) comma-separated list of allowed floats"
      },
      {
        "name": "allowInts",
        "default": "",
        "description": "allowInts"
      },
      {
        "name": "allowStrs",
        "default": "",
        "description": "(string) comma-separated list of allowed string literals"
      },
      {
        "name": "maxLitCount",
        "default": "",
        "description": "(string) maximum number of instances of a string literal that are tolerated before warn."
      }
    ],
    "timeToFix": 0
  }
]
```

To give a more detailed explanation about the issue, you should define the ```<PATTERN-ID>.md```. Example:

```
_Description_: Suggests using constant for [magic numbers](https://en.wikipedia.org/wiki/Magic_number_(programming)#Unnamed_numerical_constants) and string literals.

_Configuration_:

* `maxLitCount` : (string) maximum number of instances of a string literal that are tolerated before warn.
* `allowStr`: (string) comma-separated list of allowed string literals
* `allowInts`: (string) comma-separated list of allowed integers
* `allowFloats`: (string) comma-separated list of allowed floats

Example:

[rule.add-constant]
  arguments = [{maxLitCount = "3",allowStrs ="\"\"",allowInts="0,1,2",allowFloats="0.0,0.,1.0,1.,2.0,2."}]
```

This documentation should be generated automatically by using the Documentation Generator tool:

```
make docgeneration
```

#### Test

Follow the instructions at [codacy-plugins-test](https://github.com/codacy/codacy-plugins-test).


## What is Codacy

[Codacy](https://www.codacy.com/) is an Automated Code Review Tool that monitors
your technical debt, helps you improve your code quality, teaches best practices
to your developers, and helps you save time in Code Reviews.

### Among Codacy’s features

* Identify new Static Analysis issues
* Commit and Pull Request Analysis with GitHub, BitBucket/Stash, GitLab (and
  also direct git repositories)
* Auto-comments on Commits and Pull Requests
* Integrations with Slack, HipChat, Jira, YouTrack
* Track issues in Code Style, Security, Error Proneness, Performance, Unused
  Code and other categories

Codacy also helps keep track of Code Coverage, Code Duplication, and Code
Complexity.

Codacy supports PHP, Python, Ruby, Java, JavaScript, and Scala, among others.

### Free for Open Source

Codacy is free for Open Source projects.
