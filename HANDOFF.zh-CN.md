# Blossom Router 交接记录

[English](HANDOFF.md) | [简体中文](HANDOFF.zh-CN.md)

## 当前状态

Blossom Router 当前处于初始 MVP 阶段。Go CLI 构建为 `bloom`，支持 `local`、`ask`、`code`、`strong` 和 `auto`。Provider 使用可配置的 OpenAI-compatible chat-completions endpoint。自动路由是确定性的，`auto --dry-run` 只报告路由，不调用模型。

仓库目标位置为 `keyscome/blossom-router`。当前尚未创建正式 Release 或版本 tag。

## 已交付行为

| 区域 | 当前行为 |
| --- | --- |
| CLI | 支持参数或 stdin prompt；结果写 stdout，路由与错误信息写 stderr。 |
| 本地 | 默认 Ollama `http://localhost:11434/v1`；模型可配置。 |
| 云端 | `cheap`、`normal`、`strong`、`code` 均可配置 base URL、model 和 API key。 |
| 命令 | `local` → local，`ask` → normal，`code` → code，`strong` → strong。 |
| 自动路由 | 短通用任务 → local；批处理/中等长度 → cheap；编码 → normal；复杂/超长 → strong。 |
| 安全 | 不编译写死云端模型，不做隐藏回退、重试、串联、存储或遥测。 |
| 测试 | 路由规则和 OpenAI-compatible HTTP 路径已有单元测试。 |
| 项目展示 | 中英双语 README、项目描述、可复用主视觉和中英文发布海报。 |

## 复现维护工作流

```bash
go test ./...
go vet ./...
make build
./bin/bloom auto --dry-run "design a migration plan"
```

本地实时冒烟测试需启动 Ollama，并使用已经安装的模型：

```bash
BLOSSOM_LOCAL_MODEL=qwen2.5-coder:14b ./bin/bloom local "Reply with exactly: OK"
```

这里的模型名只是开发机器上的示例，不是产品默认值或要求。

## 事实来源

| 契约 | 位置 |
| --- | --- |
| CLI 连接和输出 | `internal/app/app.go` |
| 配置路径和环境变量覆盖 | `internal/config/config.go` |
| Provider 协议 | `internal/provider/openai.go` |
| 自动路由策略 | `internal/router/router.go` |
| 用户安装与示例 | `README.md` |
| Agent 操作规则 | `AGENTS.zh-CN.md` |

本地 YAML、Shell 环境、已安装的 Ollama 模型和生成的二进制属于机器状态，不是仓库事实。

## 已知限制

- 响应不支持流式输出。
- 只实现 chat-completions-compatible endpoint。
- 自动路由基于关键词和长度，不是语义分类。
- 没有 timeout 参数、重试、回退、成本统计或对话历史。
- CLI 解析要求 flag 位于 prompt 参数之前。
- 配置加载和 CLI 行为仍需要更完整的单元测试。

这些限制符合 MVP，除非真实使用形成更明确的下一项需求。

## 下一项安全工作

先在日常本地和云端工作流中使用本工具，再扩展范围。近期最有价值的工程工作是配置和 CLI 表驱动测试、请求超时配置，以及打包/CI。安装器现在把 `bloom` 放入 `~/.local/bin`；已经打开的 shell 可能仍需执行 `source ~/.zshrc && rehash`。不要仅为预判未来而加入 RAG、MCP、daemon 或 Agent 框架。
