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

## Agent Playbook: Updating This Repository End-to-End

This section is written for an AI coding agent (or a human) tasked with updating this repo — most commonly bumping the wrapped `revive` version, but also base image / orb / Go-dependency bumps. Follow it top to bottom; it tells you what to change, how to regenerate derived files, how to test locally, and how to interpret CI so you can iterate on failures without guessing.

### 1. What this repository is

This is a **Codacy engine**: a small Go program (built on `github.com/codacy/codacy-engine-golang-seed/v6`, see `go.mod`, `main.go`, `implementation.go`) that packages [revive](https://github.com/mgechev/revive) — a Go linter — as a Docker image Codacy's platform can run against a customer's source code. The wrapped version is pinned in a single plain-text file, `.tool_version` (currently `1.12.0`), which both `Dockerfile` and `.circleci/config.yml` read from.

The `docs/` directory is not just documentation — it is **machine-consumed configuration**:

- `docs/patterns.json` — the full list of revive rules ("patterns") Codacy knows about, their category/level, and whether each is enabled by default. Generated file, do not hand-edit.
- `docs/description/description.json` + `docs/description/*.md` — human-readable titles/descriptions per pattern, used in the Codacy UI. Generated files, do not hand-edit.
- `docs/multiple-tests/*` and `test/data/*` — fixtures used by `codacy-plugins-test` and by the Go unit tests to validate the engine actually produces the results it claims to for real code samples.
- `docs/tool-description.md` — short blurb about the tool, hand-maintained.

Both generated artifacts above come from **`docgenerator`** (`docgenerator/main.go`, `docgenerator/repository.go`, `docgenerator/defaults.go`), run via `make build-docs`. It reads `.tool_version`, downloads `RULES_DESCRIPTIONS.md` and `defaults.toml` directly from the `mgechev/revive` GitHub repo at tag `v<toolVersion>`, converts the markdown to HTML with **pandoc**, and scrapes it to build the pattern list and descriptions. This means the generator needs **network access** and **pandoc** installed locally (`brew install pandoc` on macOS).

Separately, `toolparameters/parameters.go` hand-maintains the list of configurable rule parameters (name, type, default, description) that Codacy exposes for parameterized revive rules — this file is **not** regenerated by `docgenerator` and must be edited by hand when a new/changed rule introduces or renames parameters.

### 2. Files that encode versions — check all of these on every update

| File | What it controls | What to check |
|---|---|---|
| `.tool_version` | Which revive release is installed/bundled | Bump to the target version (e.g. `1.12.0` → `1.13.0`); confirm a matching upstream tag (`v<version>`) exists on `mgechev/revive`. |
| `Dockerfile` → `go install github.com/mgechev/revive@v${TOOL_VERSION}` | Installs revive into the image at the version passed via `--build-arg TOOL_VERSION` (sourced from `.tool_version` via the Makefile) | No direct edit needed if only bumping `.tool_version`, but confirm the `golang:1.25-alpine3.22` / `alpine:3.22` base images are still current if doing a base-image bump — note the builder's Go version must satisfy the wrapped revive release's own `go.mod` minimum (e.g. revive v1.15.0 requires Go ≥ 1.25). |
| `.circleci/config.yml` → `codacy: codacy/base@12.1.5` orb | Shared CircleCI checkout/publish/tag steps | Check the latest published version if doing an infra bump. |
| `.circleci/config.yml` → `codacy_plugins_test: codacy/plugins-test@2.0.11` orb | Runs `codacy-plugins-test` (`plugins_test` job) in CI | Same as above. |
| `.circleci/config.yml` → `lint` job's own `go install github.com/mgechev/revive@v$(cat .tool_version)` | CI lints this repo's own Go code using the same revive version being bumped | Automatically picks up `.tool_version`; no separate edit needed. |
| `go.mod` / `go.sum` | Go toolchain and library versions (`codacy-engine-golang-seed`, `goquery`, `go-toml`, `testify`, etc.) | Only bump if the task scope includes a dependency update; use `go get -u -t && go mod tidy`. |

### 3. Step-by-step update procedure

1. **Bump `.tool_version`** to the target revive release.
2. **Regenerate the docs**: `make build-docs` (requires network access and pandoc). Review the diff to `docs/patterns.json` and `docs/description/*` for new/removed/renamed rules.
3. **Update `toolparameters/parameters.go` by hand** if the regenerated docs show new or changed rule parameters — this file is not auto-generated.
4. **Update Go code if revive's output/API changed** — check `command.go` (parses revive's JSON output into Codacy issues), `configuration.go` (maps Codacy parameters to revive config), and `implementation.go` (invokes the `revive` binary). Past bumps (e.g. commit `d797156`, "bump gorevive 1.12.0") touched these files alongside the doc regeneration, so do not assume a bump is docs-only.
5. **Run Go unit tests**: `make unittest` (or `go test ./...`).
6. **Build the binary**: `make build`.
7. **Build the Docker image**: `make docker` (reads `.tool_version` via the Makefile and passes it as `--build-arg TOOL_VERSION`).
8. **Run `codacy-plugins-test` locally** before pushing — clone https://github.com/codacy/codacy-plugins-test and run the relevant test commands against your local image tag (`codacy-gorevive:latest`), using the fixtures under `docs/multiple-tests/`.
9. **Iterate on failures**, re-running only the relevant test command after each fix.
10. **Commit** the version bump together with the regenerated docs files and any hand-edited Go/parameter files in one change.
11. **Push and open a PR.**
12. **Poll the PR's real CI checks until they all pass — local validation is NOT the finish line.** This repo's CI is on **CircleCI** (see the badge in this README and `.circleci/config.yml`), not GitHub Actions — the GitHub Actions workflows in `.github/workflows/` only handle Jira issue syncing and do not run tests. After every push, check the CircleCI pipeline status for your branch/PR (via the CircleCI UI, API, or `gh pr checks <pr-url>` if GitHub status checks are mirrored) and keep re-polling while any job is running until all of `unit_test`, `lint`, `publish_docker_local`, and `plugins_test` finish. If a job fails, fetch its actual log (don't guess), find the true root cause, fix it, push again (never `--no-verify`, never force-push), and re-poll. Repeat until every check is green. **The CI environment's toolchain can differ from your local one**, so a clean local run does not guarantee CI passes. Only stop iterating when every check passes, or you hit a genuine product/infra decision that needs a human.

### 4. Common failure modes and fixes

| Symptom | Cause | Fix |
|---|---|---|
| `missing go.sum entry for module providing package` | `go.sum` out of sync after a dependency bump | `go mod tidy` |
| Build fails on `zsyscall_darwin_arm64` (Apple Silicon) | Stale `golang.org/x/sys` | `go get -u golang.org/x/sys` |
| `"pandoc": executable file not found in $PATH` when running `make build-docs` | pandoc not installed locally | `brew install pandoc` (or your OS's package manager) |

### 5. Definition of done

- `.tool_version` reflects the target revive version.
- `docs/patterns.json` and `docs/description/*` regenerated via `make build-docs` and committed.
- `toolparameters/parameters.go` updated by hand if the rule set's parameters changed.
- Any Go source changes needed for compatibility with the new revive output/behavior are made and covered by unit tests.
- `make unittest`, `make build`, and `make docker` all succeed locally.
- `codacy-plugins-test` passes locally against the freshly built image.
- **After pushing and opening/updating the PR, every CI check on it (CircleCI: `unit_test`, `lint`, `publish_docker_local`, `plugins_test`, and — on `master` only — `publish_docker`/`tag_version`) is green.** Iterate on any failure until all pass.

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
