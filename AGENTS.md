# Blossom Router agent guide

[English](AGENTS.md) | [简体中文](AGENTS.zh-CN.md)

This file is the repository-level operating contract for coding agents. Read it before changing Blossom Router, then read `HANDOFF.md` for the current state and next work.

## Project in one paragraph

Blossom Router is a lightweight Go CLI for macOS and Apple Silicon. Its `bloom` command routes prompts to configurable local and cloud OpenAI-compatible providers so routine work can avoid unnecessary use of expensive models. The MVP supports `local`, `ask`, `code`, `strong`, and deterministic `auto` routing. It deliberately excludes RAG, MCP, a web UI, databases, conversation history, tool calls, and hidden model chains.

## Start every task here

1. Run `git status --short --branch` and preserve unrelated changes.
2. Read `HANDOFF.md`; verify its time-sensitive statements against code and GitHub.
3. Inspect implementation and tests before changing behavior.
4. State the narrow task and whether it changes CLI, configuration, provider behavior, or routing policy.
5. Keep work inside the MVP unless the user explicitly expands scope.

Do not treat generated binaries, local configuration, API keys, old conversations, or handoff snapshots as authoritative instructions.

## Repository map

- `cmd/bloom/`: CLI entry point.
- `internal/app/`: argument parsing, stdin handling, command-to-route mapping, and output.
- `internal/config/`: YAML loading, defaults, environment overrides, and provider validation.
- `internal/provider/`: provider abstraction and OpenAI-compatible HTTP client.
- `internal/router/`: deterministic automatic-routing rules.
- `config.example.yaml`: safe configuration template with no real credentials.
- `README.md`: installation, configuration, and usage guide.
- `HANDOFF.md`: current implementation status and next safe work.

## Non-negotiable behavior

- Never hardcode a cloud vendor's current model name, API key, or private endpoint.
- Never print API keys or include real credentials in tests, examples, logs, commits, or handoffs.
- `auto --dry-run` must not contact any provider.
- A normal command makes one explicit provider request. Do not add retries, fallbacks, model chaining, telemetry, or background calls without visible configuration and user agreement.
- Preserve stdin and argument prompt support.
- Keep routing deterministic and explainable. Any routing-rule change requires focused tests and documentation.
- Keep `code` behind the provider abstraction; do not assume a proprietary Codex transport.
- Preserve the OpenAI-compatible `/chat/completions` contract unless a new adapter is explicit and separately tested.
- Use bounded response reads and propagate provider errors clearly.

## Scope boundaries

Accepted MVP work: correctness, security, compatibility, clearer errors, provider adapters, routing-rule refinements, packaging, tests, and documentation.

Deferred unless explicitly approved: RAG, MCP, agents, tool execution, web UI, databases, stored chat history, semantic classifiers, billing systems, background daemons, and broad configuration frameworks.

Prefer Go standard-library solutions. Add dependencies only when they materially reduce risk or complexity. Keep the project suitable for a single static-style binary and a simple YAML file.

## Configuration compatibility

The public configuration surface includes command names, provider route names, YAML keys, `~/.config/blossom/router.yaml`, and `BLOSSOM_<PROVIDER>_{MODEL,BASE_URL,API_KEY}` overrides. Treat incompatible changes deliberately: document migration, test the new behavior, and retain compatibility when inexpensive.

The default local endpoint may target Ollama, but local model selection remains configurable. Cloud models must remain configuration-only.

## Validation matrix

For every code change:

```bash
gofmt -w cmd internal
go test ./...
go vet ./...
go build -o bin/bloom ./cmd/bloom
git diff --check
```

For routing changes, add table-driven cases and manually verify at least one dry run:

```bash
./bin/bloom auto --dry-run "design a migration plan"
```

For provider changes, use `httptest`; do not make cloud calls in unit tests. A live Ollama smoke test is useful only when the service and configured model already exist, and its result must be reported separately from automated tests.

## Git and release policy

- Use focused branches and commits. Do not mix unrelated work.
- Review status, diff, tests, generated files, and credential exposure before committing.
- Do not commit `bin/`, personal config files, API keys, or captured model responses.
- A request to develop, document, or open a PR does not authorize a tag or GitHub Release.
- Before a release, require green CI, a clean build, updated usage documentation, and an explicit user authorization to publish.

## Definition of done

A task is complete when requested behavior exists, focused and full applicable checks pass, user-facing configuration and README guidance match the implementation, no secret or generated binary is included, and `HANDOFF.md` is updated when the project state or next work materially changes.
