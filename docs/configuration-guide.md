# Effective configuration guide

[English](configuration-guide.md) | [简体中文](configuration-guide.zh-CN.md)

## Recommended first setup

Use four auto-routing slots and one optional explicit coding slot:

- `local`: a reliable Ollama model that fits memory.
- `cheap`: low-cost cloud work such as translation and summaries.
- `normal`: balanced general and coding work.
- `strong`: difficult reasoning, architecture, security, and root-cause analysis.
- `code`: optional explicit coding endpoint; `auto` does not choose it.

If you do not yet have four distinct models, point cheap and normal to the same endpoint. Completeness is more important than artificial differentiation because `auto` returns a clear error rather than silently falling back.

## Example

```yaml
providers:
  local:
    base_url: http://localhost:11434/v1
    model: qwen2.5-coder:14b
  cheap:
    base_url: https://your-provider.example/v1
    model: your-cheap-model
    api_key_env: BLOSSOM_CHEAP_API_KEY
  normal:
    base_url: https://your-provider.example/v1
    model: your-normal-model
    api_key_env: BLOSSOM_NORMAL_API_KEY
  strong:
    base_url: https://your-provider.example/v1
    model: your-strong-model
    api_key_env: BLOSSOM_STRONG_API_KEY
  code:
    base_url: https://your-code-provider.example/v1
    model: your-code-model
    api_key_env: BLOSSOM_CODE_API_KEY
```

Keep keys outside YAML:

```bash
export BLOSSOM_CHEAP_API_KEY=...
export BLOSSOM_NORMAL_API_KEY=...
export BLOSSOM_STRONG_API_KEY=...
export BLOSSOM_CODE_API_KEY=...
```

Start the lower-complexity interface with `bloom serve`. It reads the same config and never returns resolved keys to the browser.
