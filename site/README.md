# Blossom Router site

[English](README.md) | [简体中文](README.zh-CN.md)

This directory contains the bilingual public project site for Blossom Router. English is served at `/`; Simplified Chinese is served at `/zh`.

The public site explains the routing model, provides a local-only fit assessment and YAML configuration generator, answers common questions, and points users to the localhost operating UI. It never accepts API keys or calls configured AI providers.

## Develop and validate

Use Node.js 22.13 or newer:

```bash
npm install
npm run dev
npm test
npm run lint
```

The project uses Vinext and the Sites build contract. Hosting identity is stored in `.openai/hosting.json`; secrets and runtime credentials must never be stored there.
