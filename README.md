# plugin-template

[![Latest Release](https://img.shields.io/github/v/release/SemRels/plugin-template?label=version&color=blue)](https://github.com/SemRels/plugin-template/releases/latest)
[![CI](https://github.com/SemRels/plugin-template/actions/workflows/ci.yml/badge.svg)](https://github.com/SemRels/plugin-template/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/SemRels/plugin-template)](LICENSE)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/SemRels/plugin-template/badge)](https://scorecard.dev/viewer/?uri=github.com/SemRels/plugin-template)
[![REUSE status](https://api.reuse.software/badge/github.com/SemRels/plugin-template)](https://api.reuse.software/info/github.com/SemRels/plugin-template)

Template repository for SemRels plugins built as standalone Go subprocess binaries.

This template reflects the current SemRels plugin architecture: semrel executes a plain Go CLI binary as a subprocess, passes plugin configuration through `SEMREL_PLUGIN_*` variables, passes release context through `SEMREL_*` variables, reads plugin output from standard output, and treats exit code `0` as success and any non-zero exit code as failure. There is no gRPC transport, no network listener, and no protobuf plumbing to wire up.

## What the example plugin does

The scaffolded example is intentionally phase-agnostic. It reads `SEMREL_VERSION` and `SEMREL_DRY_RUN`, writes the required schema handshake to standard error, and prints a confirmation message to standard output:

```text
stderr: plugin_schema_version=1
stdout: example plugin invoked for version 1.2.3 (dry-run: false)
```

## Plugin contract

| Item | Behavior |
| --- | --- |
| `plugin_schema_version=1` | The **first line written to stderr** must be `plugin_schema_version=1`. |
| `SEMREL_VERSION` | Required release version for this example plugin. |
| `SEMREL_DRY_RUN` | Optional boolean (`true`/`false`) indicating dry-run mode. |
| `SEMREL_PLUGIN_*` | Reserved for your phase-specific plugin configuration. |
| stdout | Write the plugin's functional output here. |
| stderr | Write diagnostics, validation errors, and logs here. |
| exit code | `0` means success; any non-zero code means failure. |

## Repository layout

```text
cmd/plugin/              Thin environment-variable wrapper and process entry point
internal/plugin/         Pure, testable plugin logic with no os/env dependencies
.github/workflows/       CI, release, security, template-sync, and semrel self-release automation
.semrel.yaml             semrel dogfooding configuration for this repository
```

## Use this template for a real plugin

1. Create a new repository from this template or copy it into your plugin repository.
2. Rename the Go module in `go.mod` to `github.com/SemRels/<your-plugin-repo>`.
3. Rename the binary/package descriptions from `plugin-template` to your plugin name.
4. Implement your phase-specific logic in `internal/plugin/` as pure Go.
5. Keep `cmd/plugin/main.go` thin: read `SEMREL_*` / `SEMREL_PLUGIN_*` with `os.Getenv` (or an injected `getenv` in tests), translate them into a config struct, then call the internal package.
6. Extend the README tables with every `SEMREL_PLUGIN_*` and `SEMREL_*` variable your plugin consumes.
7. Update `.semrel.yaml` and replace the example phase assignment comments/configuration with the plugins and phases your real plugin release process needs.

## Local development

```bash
go build ./...
go test ./...
```

## Example semrel configuration

```yaml
plugins:
  - name: example-plugin
    path: ~/.semrel/plugins/semrel-plugin-example
    env:
      SEMREL_PLUGIN_SAMPLE_OPTION: "value"
```

## License

Apache-2.0
