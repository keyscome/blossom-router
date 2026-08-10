# Blossom Router handoff

[English](HANDOFF.md) | [简体中文](HANDOFF.zh-CN.md)

## Current status

Blossom Router is at its initial MVP. The Go CLI builds as `bloom` and supports `local`, `ask`, `code`, `strong`, and `auto`. Providers use a configurable OpenAI-compatible chat-completions endpoint. Automatic routing is deterministic, and `auto --dry-run` reports a route without calling a model.

The repository is intended to live at `keyscome/blossom-router`. No formal release or version tag exists yet.

## Delivered behavior

| Area | Current behavior |
| --- | --- |
| CLI | Prompt from arguments or stdin; clear stdout results and stderr routing/error messages. |
| Local | Defaults to Ollama at `http://localhost:11434/v1`; model is configurable. |
| Cloud | `cheap`, `normal`, `strong`, and `code` slots use configurable base URL, model, and API key. |
| Commands | `local` → local, `ask` → normal, `code` → code, `strong` → strong. |
| Auto routing | Short general → local; batch/medium → cheap; coding → normal; complex/very long → strong. |
| Safety | No compiled cloud model names, hidden fallbacks, retries, chains, storage, or telemetry. |
| Tests | Routing rules and the OpenAI-compatible HTTP path have unit coverage. |
| Presentation | Bilingual README, project descriptions, a reusable hero image, and English/Chinese launch posters. |

## Reproduce the maintained workflow

```bash
go test ./...
go vet ./...
make build
./bin/bloom auto --dry-run "design a migration plan"
```

For a live local smoke test, start Ollama and use a model already installed:

```bash
BLOSSOM_LOCAL_MODEL=qwen2.5-coder:14b ./bin/bloom local "Reply with exactly: OK"
```

The model name above is an example from the development machine, not a product default or requirement.

## Sources of truth

| Contract | Location |
| --- | --- |
| CLI wiring and output | `internal/app/app.go` |
| Config paths and environment overrides | `internal/config/config.go` |
| Provider protocol | `internal/provider/openai.go` |
| Automatic routing policy | `internal/router/router.go` |
| User installation and examples | `README.md` |
| Agent operating rules | `AGENTS.md` |

Local YAML files, shell environment, installed Ollama models, and generated binaries are machine state, not repository truth.

## Known limitations

- Responses are non-streaming.
- Only chat-completions-compatible endpoints are implemented.
- Automatic routing is keyword/length based, not semantic classification.
- There is no timeout flag, retry policy, fallback, cost accounting, or conversation history.
- CLI parsing expects flags before prompt arguments.
- Configuration loading and CLI behavior need broader unit coverage.

These limitations are consistent with the MVP unless real usage establishes a narrower next requirement.

## Next safe work

Use the tool in normal local/cloud workflows before expanding scope. The most useful near-term engineering work is configuration and CLI table-driven testing, request timeout configuration, and packaging/CI. The installer now places `bloom` in `~/.local/bin`; an already-open shell may still need `source ~/.zshrc && rehash`. Do not add RAG, MCP, a daemon, or an agent framework merely to anticipate future use.
