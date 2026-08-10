# Blossom Router FAQ

[English](FAQ.md) | [简体中文](FAQ.zh-CN.md)

## What problem does Blossom Router solve?

It reduces unnecessary premium-model usage when your workload mixes routine prompts with genuinely difficult tasks. One CLI maps explicit commands or deterministic rules to configurable local, cheap, normal, strong, and code provider slots.

## Is it useful if I use only one model?

Usually not yet. Calling that provider directly is simpler. Blossom Router becomes useful when you use at least two tiers, want a stable command across providers, or need to protect premium quota from repetitive work.

## Does automatic routing call another model?

No. It uses local keyword and prompt-length rules. `bloom auto --dry-run` reports the route and reason without making any provider request.

## Which routes must be configured for `auto`?

Configure `local`, `cheap`, `normal`, and `strong`, because the current deterministic rules may select any of them. `code` is explicit and is not selected by `auto`. The MVP does not silently fall back when a selected route is missing.

## Can multiple routes use the same provider?

Yes. They can use different models from one vendor, different vendors, or even the same endpoint and model while you validate the workflow. The route names describe intent and cost tier, not a required company.

## Where should API keys go?

Use `api_key_env` in YAML and export the named environment variable. Do not put a real key in a committed config file. The public website never asks for a key, and the local UI reads the resolved server-side configuration without returning credentials to the browser.

## Does the browser UI expose my local service?

`bloom serve` accepts only loopback addresses such as `127.0.0.1`. It refuses `0.0.0.0` because the UI can call your configured providers. It opens at `http://127.0.0.1:7331` by default and stops with Ctrl+C.

## Can the `code` route use Codex?

It can use a configurable OpenAI-compatible chat-completions endpoint. The MVP does not assume a proprietary Codex transport or add hidden agent calls.

## Which Ollama model should I choose?

Start with a model already stable on your Mac. For the current development machine, `qwen2.5-coder:14b` is installed and works as a practical local coding/general model. Model availability and memory requirements vary, so keep this choice configurable.

## How do I know whether routing is saving money?

Sample real tasks for one week and record route, completion time, whether you accepted the answer, and whether you manually escalated. Useful signals are local/cheap share, escalation rate, failure rate, latency, and estimated cloud spend avoided. See the [evaluation guide](docs/evaluation.md).

## Why are retries and fallback absent?

They can create invisible cost and make routing harder to understand. The MVP makes one visible request and returns the error. Any future retry or fallback policy should be explicit, bounded, configurable, and measurable.

## Does the public site operate my models?

No. It provides documentation, a fit check, and a configuration generator. Use `bloom serve` for the local operating interface.
