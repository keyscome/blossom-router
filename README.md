# Doudou CLI (`dd`)

A small macOS-friendly CLI that routes everyday AI work to local or configurable cloud models. It is intentionally limited to a single binary, one YAML file, one OpenAI-compatible protocol, and simple deterministic routing.

## Install

Requirements: Go 1.24+; Ollama is optional but recommended for local use.

```bash
make test
make build
install -m 755 bin/dd ~/.local/bin/dd
mkdir -p ~/.config/doudou
cp config.example.yaml ~/.config/doudou/config.yaml
```

Ensure `~/.local/bin` is on `PATH`. Alternatively, `make install` installs through Go into `$GOBIN` (or `$GOPATH/bin`).

## Configure

Edit `~/.config/doudou/config.yaml`. Every route points to an OpenAI-compatible endpoint; model names are never compiled into cloud commands. API keys can be provided indirectly with `api_key_env`, or directly with `api_key` (not recommended).

Each configured provider also supports environment overrides:

```bash
export DOUDOU_LOCAL_MODEL=qwen3:8b
export DOUDOU_LOCAL_BASE_URL=http://localhost:11434/v1
export DOUDOU_STRONG_MODEL=your-model
export DOUDOU_STRONG_BASE_URL=https://api.example.com/v1
export DOUDOU_STRONG_API_KEY=secret
```

Environment override names use `DOUDOU_<PROVIDER>_{MODEL,BASE_URL,API_KEY}`. Values from environment variables override YAML.

## Use

```bash
dd local "summarize this log"
dd ask "explain this deployment failure"
git diff | dd code "review this diff"
dd strong "compare these architecture options"
dd auto "translate this release note"
dd auto --dry-run "design a migration plan"
cat error.log | dd auto
```

`auto` uses transparent rules: short general prompts go local; summary/translation/classification and medium prompts go cheap; coding keywords go normal; architecture/security/migration/root-cause keywords and prompts over 3,000 characters go strong. It prints the chosen route to stderr. `--dry-run` never calls a model.

`code` is a separate configurable provider slot. It can target any Codex/OpenAI-compatible chat-completions endpoint, but it makes exactly one request and performs no hidden planning, retry, or model chaining.

## Ollama and Qwen

Install and start Ollama, then pull the model named in your config:

```bash
brew install ollama
ollama serve
ollama pull qwen3:8b
dd local "用三句话总结 Kubernetes readiness probe"
```

Choose a Qwen size that fits your Mac memory and change only `providers.local.model`. Ollama exposes the compatible endpoint at `http://localhost:11434/v1`; it does not require an API key.

## Cloud endpoints

For each cloud tier, set its `base_url`, `model`, and `api_key_env`. Different tiers may use the same vendor with different models, or entirely different compatible vendors. No particular OpenAI model is assumed. If an endpoint does not implement `/v1/chat/completions`, place a compatible gateway in front of it or extend the `Provider` interface in `internal/provider`.

## Scope

This MVP deliberately excludes RAG, MCP, a web UI, databases, conversation history, streaming, tool calls, and automatic fallback. Those can be added later only when real usage justifies them.
