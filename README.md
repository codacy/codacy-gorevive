# Codacy Revive

Codacy tool to add support for [revive](https://github.com/mgechev/revive).

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/eac3dec12b244a1390f1f414246d0d6d)](https://app.codacy.com/gh/codacy/codacy-gorevive/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![CircleCI](https://circleci.com/gh/codacy/codacy-gorevive.svg?style=svg)](https://circleci.com/gh/codacy/codacy-gorevive)

# Usage

You can create the docker by doing:

```bash
make docker
```

The docker is ran with the following command:

```bash
docker run -it -v $srcDir:/src <DOCKER_NAME>:<DOCKER_VERSION>
```

To run the tool using a custom configuration file, run docker with the following command:

```bash
docker run -it -v $srcDir:/src -v $codacyrcConfig:/.codacyrc <DOCKER_NAME>:<DOCKER_VERSION>
```

## Tool Developer Documentation

[Tool Developer Guide](https://support.codacy.com/hc/en-us/articles/207994725-Tool-Developer-Guide)


[Tool Developer Guide - Using Scala](https://support.codacy.com/hc/en-us/articles/207280379-Tool-Developer-Guide-Using-Scala)


## Structure

-   To run the tool we provide the configuration file, `/.codacyrc`, with the language to run and optional parameters your tool might need.
-   The source code to be analysed will be located in `/src`, meaning that when provided in the configuration, the file paths are relative to `/src`.

Check more information [here](https://github.com/codacy/codacy-example-tool#structure)

#### Tool documentation

At Codacy we strive to provide the best value to our users and, to accomplish that, we document our patterns so that the user can better understand the problem and fix it.

The documentation for the tool must always be updated before submitting the docker.

To get more information on the tool documentation, check [here](https://github.com/codacy/codacy-example-tool#tool-documentation)

#### Generate documentation

This documentation should be generated automatically by using the Documentation Generator tool:

```bash
make build-docs
```

#### Bumping the version
-   update the .tool_version file with the latest revive version
-   upgrade dependencies
 ```bash
  go get -u -t
  go install
  go fix
  go mod tidy
 ```
-   generate documentation
-   if rules change add/edit related parameters in `toolparameters/parameter.go`.

##### Common errors
###### missing go.sum entry for module providing package
```bash
go mod tidy
```
###### zsyscall_darwin_arm64 (Apple m1)
```bash
go get -u golang.org/x/sys
```
###### "pandoc": executable file not found in $PATH
```bash
brew install pandoc
```

#### Test

Follow the instructions at [codacy-plugins-test](https://github.com/codacy/codacy-plugins-test).


## What is Codacy

[Codacy](https://www.codacy.com/) is an Automated Code Review Tool that monitors
your technical debt, helps you improve your code quality, teaches best practices
to your developers, and helps you save time in Code Reviews.

### Among Codacy’s features

-   Identify new Static Analysis issues
-   Commit and Pull Request Analysis with GitHub, BitBucket/Stash, GitLab (and also direct git repositories)
-   Auto-comments on Commits and Pull Requests
-   Integrations with Slack, HipChat, Jira, YouTrack
-   Track issues in Code Style, Security, Error Proneness, Performance, Unused Code and other categories

Codacy also helps keep track of Code Coverage, Code Duplication, and Code
Complexity.

Codacy supports PHP, Python, Ruby, Java, JavaScript, and Scala, among others.

### Free for Open Source

Codacy is free for Open Source projects.
