# Practical evaluation guide

[English](evaluation.md) | [简体中文](evaluation.zh-CN.md)

## When the project is useful

Blossom Router is a strong fit when you use local and cloud models, handle many small repeatable prompts, and want premium usage to be an explicit decision. It is a weak fit when one provider handles everything, your main need is a fully autonomous tool-using agent, or maintaining endpoint configuration costs more time than routing saves.

## One-week evaluation

Use 30–50 representative tasks. For each task, record:

| Field | Why it matters |
| --- | --- |
| Intended task type | Confirms the sample reflects real work. |
| Selected route | Shows workload distribution. |
| Accepted without escalation | Measures route quality. |
| Manual escalation | Detects under-routing. |
| Completion time | Detects local models that are too slow. |
| Estimated cloud cost | Measures actual savings rather than call count. |

Recommended decision signals:

- Local + cheap share of at least 50% indicates meaningful quota offload.
- Manual escalation below 15% suggests the simple router is sufficiently accurate.
- Repeated failures in one task category indicate a routing-rule or model-fit problem.
- If operating the router takes longer than the quota it saves, simplify to fewer tiers.

These are operational starting points, not universal claims. Declare your own thresholds before reviewing the week of results.

## Suggested rollout

1. Configure every route used by `auto`, even if cheap and normal initially share one endpoint.
2. Use `--dry-run` for the first 10–20 tasks and compare the suggested route with your judgment.
3. Start real calls after the rules look reasonable.
4. Review the sample after one week; change one rule or model at a time.
5. Keep strong work explicit until automatic routing has earned trust.
