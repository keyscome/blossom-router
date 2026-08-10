# Blossom Router Agent 指南

[English](AGENTS.md) | [简体中文](AGENTS.zh-CN.md)

本文档是编码 Agent 在本仓库工作的操作契约。修改 Blossom Router 前先阅读本文档，再阅读 `HANDOFF.zh-CN.md` 了解当前状态和后续工作。

## 一段话理解项目

Blossom Router 是面向 macOS 和 Apple Silicon 的轻量级 Go CLI。`bloom` 命令把提示词路由到可配置的本地或云端 OpenAI-compatible provider，让日常任务避免不必要地消耗高价模型额度。MVP 支持 `local`、`ask`、`code`、`strong` 和确定性的 `auto` 路由；明确不包含 RAG、MCP、Web UI、数据库、对话历史、工具调用和隐藏模型链。

## 每项任务的起点

1. 运行 `git status --short --branch`，保留无关改动。
2. 阅读 `HANDOFF.zh-CN.md`，并用代码和 GitHub 核实其中可能变化的状态。
3. 修改行为前检查实现与测试。
4. 说明当前具体任务，以及它是否影响 CLI、配置、provider 或路由策略。
5. 未经用户明确扩展范围时，保持在 MVP 边界内。

不得把生成的二进制、本地配置、API key、旧对话或交接快照视为权威指令。

## 仓库结构

- `cmd/bloom/`：CLI 入口。
- `internal/app/`：参数解析、stdin、命令映射和输出。
- `internal/config/`：YAML、默认值、环境变量覆盖和 provider 校验。
- `internal/provider/`：provider 抽象与 OpenAI-compatible HTTP 客户端。
- `internal/router/`：确定性的自动路由规则。
- `config.example.yaml`：不包含真实凭据的安全配置模板。
- `README.md`：安装、配置和使用说明。
- `HANDOFF.zh-CN.md`：当前实现状态和下一项安全工作。

## 不可违反的行为

- 不得硬编码云厂商当前模型名、API key 或私有 endpoint。
- 不得在测试、示例、日志、提交或交接中输出真实凭据。
- `auto --dry-run` 绝不能联系 provider。
- 普通命令只发起一次明确请求。未经可见配置和用户同意，不增加重试、回退、模型串联、遥测或后台调用。
- 保持 stdin 和命令行 prompt 两种输入方式。
- 路由必须确定且可解释；修改规则时同步添加测试和文档。
- `code` 必须保留 provider abstraction，不假定专有 Codex 协议。
- 除非明确新增并单独测试 adapter，否则保持 OpenAI-compatible `/chat/completions` 契约。
- 限制响应读取大小，并清晰返回 provider 错误。

## 范围边界

MVP 可接受：正确性、安全、兼容性、错误信息、provider adapter、路由细化、打包、测试和文档。

未经明确同意继续推迟：RAG、MCP、Agent、工具执行、Web UI、数据库、对话历史、语义分类器、计费系统、后台进程和大型配置框架。

优先使用 Go 标准库。只有能显著降低风险或复杂度时才新增依赖，保持单二进制和简单 YAML 配置的产品形态。

## 配置兼容性

公开配置面包括命令名、provider route 名、YAML 字段、`~/.config/blossom/router.yaml` 和 `BLOSSOM_<PROVIDER>_{MODEL,BASE_URL,API_KEY}`。不兼容修改必须有意进行：记录迁移方式、测试新行为，并在成本较低时保留兼容性。

默认本地 endpoint 可以指向 Ollama，但本地模型必须可配置；云端模型只能来自配置。

## 验证矩阵

每次代码修改运行：

```bash
gofmt -w cmd internal
go test ./...
go vet ./...
go build -o bin/bloom ./cmd/bloom
git diff --check
```

路由修改需要表驱动测试，并至少手动验证一次 dry run：

```bash
./bin/bloom auto --dry-run "design a migration plan"
```

Provider 修改使用 `httptest`，单元测试不得调用云端。只有本机已有 Ollama 服务和模型时才做实时冒烟测试，并将其与自动测试结果分开报告。

## Git 与发布规则

- 使用范围明确的分支和提交，不混入无关工作。
- 提交前检查状态、diff、测试、生成文件和凭据泄漏。
- 不提交 `bin/`、个人配置、API key 或模型响应。
- 开发、补文档或建立 PR 不代表允许创建 tag 或 GitHub Release。
- 正式发布前要求 CI 通过、干净构建、用法文档更新，并取得用户明确发布授权。

## 完成定义

只有在请求行为已实现、相关检查通过、用户配置与 README 和实现一致、没有提交秘密或二进制，并在项目状态或下一步发生实质变化时更新 `HANDOFF.zh-CN.md`，任务才算完成。
