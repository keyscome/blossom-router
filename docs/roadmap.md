# Blossom Router product roadmap

[English](roadmap.md) | [简体中文](roadmap.zh-CN.md)

Status date: 2026-08-10

## Product goal

Blossom Router should reduce unnecessary premium-model usage without becoming another complex AI platform. Its durable promise is: one small interface, explicit provider configuration, local-first routing, and no invisible model calls.

## Original requirements and current state

| Requirement | State | Release gate |
| --- | --- | --- |
| `local`, `ask`, `code`, `strong`, `auto` commands | Delivered | Preserve compatibility and tests. |
| Ollama OpenAI-compatible local provider | Delivered | Live smoke test on Apple Silicon. |
| Configurable cloud provider/model/key/base URL | Delivered | Add configuration diagnostics before 1.0. |
| Explainable automatic routing and dry-run | Delivered | Measure escalation rate on real tasks. |
| Stdin and argument prompts | Delivered | Add CLI regression coverage. |
| Go single-binary workflow | Delivered | Add reproducible release binaries. |
| Local browser operating UI | Delivered after MVP | Keep loopback-only and credential-safe. |
| Bilingual docs, FAQ, evaluation, and project site | Delivered | Maintain English/Chinese parity. |
| RAG, MCP, remote web UI, database | Explicitly deferred | Require validated demand after 1.0. |

## Priority policy

### Must

- correctness, clear errors, secret safety, loopback UI safety;
- complete configuration for every route selected by `auto`;
- deterministic tests and clean installation on macOS/Apple Silicon;
- synchronized English and Simplified Chinese user documentation.

### Should

- `bloom doctor` or equivalent configuration diagnostics;
- request timeout controls and graceful cancellation;
- streaming output without changing provider selection semantics;
- measurable one-week evaluation of route quality, latency, and quota offload;
- CI, signed release artifacts, checksums, and a simple package channel.

### Could

- explicit provider adapters beyond chat completions;
- opt-in, bounded fallback with visible cost rules;
- local-only usage export for evaluation;
- shell completion and additional packaging formats.

### Not before 1.0

- RAG, MCP, autonomous agent loops, databases, accounts, hosted secret storage,
  remote multi-user UI, or opaque semantic routing.

## Delivery schedule

| Phase | Target | Goal | Planned requirements | Exit criteria |
| --- | --- | --- | --- | --- |
| P0 Foundation | Complete | Prove the routing concept | CLI, YAML/env config, Ollama, cloud tiers, auto dry-run, local UI, bilingual site | Current main builds and core tests pass. |
| P1 Stabilization | 2026-08-10 → 2026-08-16 | Prepare `v0.1.0` | Typography/accessibility fixes, CLI/config tests, `bloom doctor`, timeout baseline, install acceptance | Clean Mac install, all route configs validated, docs parity, no known critical defect. |
| P2 Daily usefulness | 2026-08-17 → 2026-08-30 | Prepare `v0.2.0` | Streaming, cancellation, route explanations, one-week evaluation template/results | 30–50 real tasks, declared thresholds, useful quota offload without excessive escalation. |
| P3 Provider depth | 2026-08-31 → 2026-09-20 | Prepare `v0.3.0` | Explicit adapters, bounded opt-in fallback, cost guardrails, compatibility tests | No hidden calls; every extra request is visible, bounded, tested, and configurable. |
| P4 Trust and packaging | 2026-09-21 → 2026-10-18 | Prepare `v1.0.0` | Stable CLI/config contract, CI matrix, signed binaries, checksums, package channel, upgrade guide | Clean-clone acceptance, green CI, security review, bilingual release docs, explicit release approval. |

Dates are planning targets, not promises. Scope moves before dates: unfinished
quality gates delay a release rather than being waived.

## Success metrics

- at least 50% of sampled routine work handled by local + cheap routes;
- less than 15% manual escalation during the evaluation sample;
- zero hidden provider calls and zero credential exposure;
- predictable error behavior when a route is missing or unavailable;
- installation-to-first-local-answer in under 15 minutes on a prepared Mac;
- English and Chinese docs updated in the same pull request.

These thresholds are starting hypotheses. Record them before evaluation and
change them only for the next evaluation cycle, never after seeing a result.
