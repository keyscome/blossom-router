# Blossom Router release plan

[English](release-plan.md) | [简体中文](release-plan.zh-CN.md)

Status date: 2026-08-10. The repository has no formal release tag yet.

## Version sequence

| Version | Target | User value | Required scope | Explicitly excluded |
| --- | --- | --- | --- | --- |
| `v0.1.0` | 2026-08-16 | Safe public preview for local-first routing | Existing MVP, typography fix, config diagnostics, timeout baseline, regression tests, install verification | Streaming, fallback, new provider protocols |
| `v0.2.0` | 2026-08-30 | Comfortable daily terminal and local-UI use | Streaming, cancellation, clearer route explanations, evaluation evidence | Remote UI, accounts, databases |
| `v0.3.0` | 2026-09-20 | Broader compatible-provider use with cost control | Explicit adapters, compatibility suite, optional bounded fallback, cost guardrails | Opaque routing and autonomous loops |
| `v1.0.0` | 2026-10-18 | Stable contract suitable for routine adoption | Stable CLI/YAML/env contract, signed macOS artifacts, checksums, CI matrix, clean-install acceptance, upgrade guide | RAG, MCP, hosted secrets, multi-user service |

## Release authorization levels

1. **Development:** implementation and tests on a focused branch.
2. **Release preparation:** PR, changelog, version notes, clean-install test,
   artifact plan, and CI evidence.
3. **Formal release:** tag, binaries, checksums, package publication, and GitHub
   Release. This requires explicit user authorization for that version.

A merged PR is not automatically a release.

## Gate checklist

Every release requires:

- all applicable Go and site tests green on the release commit;
- `go vet`, site lint, and production site build green;
- clean installation and first-response smoke test on macOS/Apple Silicon;
- no secrets, generated local config, or unreviewed binaries in Git;
- synchronized English and Simplified Chinese README, FAQ, roadmap, changelog,
  and release notes;
- documented compatibility impact for CLI, YAML, environment variables, and
  provider protocol;
- known limitations and upstream dependency advisories stated honestly;
- explicit approval before tag, release asset publication, or package update.

## Versioning policy

- `0.x.0`: a coherent preview capability or adoption milestone.
- `0.x.y`: preview bug, compatibility, documentation, or packaging correction.
- `1.x.0`: backward-compatible user capability after the stable contract.
- `1.x.y`: stable bug, security, compatibility, or packaging correction.
- `2.0.0`: incompatible CLI, configuration, provider, or automation contract.

## Rollback and support

- Keep previous release artifacts and checksums available.
- Document configuration migrations before an incompatible change ships.
- For a broken release, stop distribution, publish a bilingual notice, and
  issue a patch rather than silently replacing an existing artifact.
- Support the latest minor release during preview; define a longer support
  window before `v1.0.0`.
