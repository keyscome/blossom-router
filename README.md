# Blossom Router (`bloom`)

[English](README.md) | [简体中文](README.zh-CN.md)

![Blossom Router — one prompt, the right model](docs/assets/blossom-router-hero.jpg)

**Route everyday AI work wisely.** Blossom Router is a small macOS-friendly CLI in the Blossom ecosystem that sends each prompt to a configurable local or cloud model tier. Routine work stays local or cheap; difficult work can be sent explicitly to stronger models. It remains intentionally limited to a single binary, one YAML file, one OpenAI-compatible protocol, and deterministic routing you can inspect before it runs.

## Why Blossom Router?

- Save premium-model quota by keeping short, repetitive work local.
- Use one command across Ollama and OpenAI-compatible cloud endpoints.
- Keep provider names, model IDs, endpoints, and credentials out of the binary.
- Preview automatic decisions with `bloom auto --dry-run` without calling a model.
- Avoid hidden retries, fallbacks, model chains, telemetry, and stored conversations.

## Install

Requirements: Go 1.24+; Ollama is optional but recommended for local use.

```bash
make test
make install
mkdir -p ~/.config/blossom
cp config.example.yaml ~/.config/blossom/router.yaml
```

`make install` writes `bloom` to `~/.local/bin` by default. If `which bloom` still reports `not found`, activate that directory in the current zsh session:

```bash
export PATH="$HOME/.local/bin:$PATH"
rehash
```

To keep it available in new terminals, add the export to `~/.zshrc`, then run `source ~/.zshrc`. Your existing terminal does not automatically reload a changed shell profile.

To install elsewhere:

```bash
make install PREFIX=/usr/local
```

## Configure

Edit `~/.config/blossom/router.yaml`. Every route points to an OpenAI-compatible endpoint; model names are never compiled into cloud commands. API keys can be provided indirectly with `api_key_env`, or directly with `api_key` (not recommended).

Each configured provider also supports environment overrides:

```bash
export BLOSSOM_LOCAL_MODEL=qwen3:8b
export BLOSSOM_LOCAL_BASE_URL=http://localhost:11434/v1
export BLOSSOM_STRONG_MODEL=your-model
export BLOSSOM_STRONG_BASE_URL=https://api.example.com/v1
export BLOSSOM_STRONG_API_KEY=secret
```

Environment override names use `BLOSSOM_<PROVIDER>_{MODEL,BASE_URL,API_KEY}`. Values from environment variables override YAML.

## Use

```bash
bloom local "summarize this log"
bloom ask "explain this deployment failure"
git diff | bloom code "review this diff"
bloom strong "compare these architecture options"
bloom auto "translate this release note"
bloom auto --dry-run "design a migration plan"
cat error.log | bloom auto
bloom serve
```

`auto` uses transparent rules: short general prompts go local; summary/translation/classification and medium prompts go cheap; coding keywords go normal; architecture/security/migration/root-cause keywords and prompts over 3,000 characters go strong. It prints the chosen route to stderr. `--dry-run` never calls a model.

`code` is a separate configurable provider slot. It can target any Codex/OpenAI-compatible chat-completions endpoint, but it makes exactly one request and performs no hidden planning, retry, or model chaining.

### Local browser UI

Run `bloom serve` to open `http://127.0.0.1:7331`. The local-only control panel lets you select a route, paste a prompt, preview automatic routing, run the request, and copy the result. It uses the same YAML and environment variables as the CLI; resolved API keys stay server-side. The server refuses non-loopback addresses.

## Ollama and Qwen

Install and start Ollama, then pull the model named in your config:

```bash
brew install ollama
ollama serve
ollama pull qwen3:8b
bloom local "用三句话总结 Kubernetes readiness probe"
```

Choose a Qwen size that fits your Mac memory and change only `providers.local.model`. Ollama exposes the compatible endpoint at `http://localhost:11434/v1`; it does not require an API key.

## Cloud endpoints

For each cloud tier, set its `base_url`, `model`, and `api_key_env`. Different tiers may use the same vendor with different models, or entirely different compatible vendors. No particular OpenAI model is assumed. If an endpoint does not implement `/v1/chat/completions`, place a compatible gateway in front of it or extend the `Provider` interface in `internal/provider`.

## Scope

This MVP deliberately excludes RAG, MCP, a web UI, databases, conversation history, streaming, tool calls, and automatic fallback. Those can be added later only when real usage justifies them.

## Project materials

- [Public project site](https://alanthssss.github.io/blossom-router)
- [FAQ](FAQ.md)
- [Practical evaluation guide](docs/evaluation.md)
- [Effective configuration guide](docs/configuration-guide.md)
- [Product roadmap](docs/roadmap.md)
- [Release plan](docs/release-plan.md)
- [Project description and messaging](docs/project-description.md)
- [English launch poster](docs/assets/blossom-router-poster-en.jpg)
- [Simplified Chinese launch poster](docs/assets/blossom-router-poster-zh-CN.jpg)
- [Agent guide](AGENTS.md) and [handoff](HANDOFF.md)
