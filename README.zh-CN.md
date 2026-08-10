# Blossom Router (`bloom`)

[English](README.md) | [简体中文](README.zh-CN.md)

![Blossom Router——一个提示词，交给合适的模型](docs/assets/blossom-router-hero.jpg)

**让每一次 AI 调用，都恰到好处。** Blossom Router 是 Blossom 生态中的轻量级 macOS CLI，把每个提示词交给可配置的本地或云端模型层级。日常工作留在本地或经济模型，困难任务再明确交给更强模型。它刻意保持简单：一个二进制、一个 YAML 文件、一套 OpenAI-compatible 协议，以及执行前可检查的确定性路由。

## 为什么使用 Blossom Router？

- 短小、重复的任务优先在本地完成，节省高价模型额度。
- 用一个命令连接 Ollama 和 OpenAI-compatible 云端 endpoint。
- Provider、模型 ID、endpoint 和凭据都不写死在二进制中。
- `bloom auto --dry-run` 不调用模型即可预览自动路由决定。
- 不进行隐藏重试、自动回退、模型串联、遥测或对话存储。

## 安装

要求 Go 1.24+；本地模式建议安装 Ollama，但不是构建所必需。

```bash
make test
make install
mkdir -p ~/.config/blossom
cp config.example.yaml ~/.config/blossom/router.yaml
```

`make install` 默认把 `bloom` 安装到 `~/.local/bin`。如果 `which bloom` 仍显示 `not found`，在当前 zsh 会话执行：

```bash
export PATH="$HOME/.local/bin:$PATH"
rehash
```

若要让新终端永久生效，把这条 `export` 加入 `~/.zshrc`，再执行 `source ~/.zshrc`。已经打开的终端不会自动重新读取修改后的配置。

也可以指定其他安装目录：

```bash
make install PREFIX=/usr/local
```

## 配置

编辑 `~/.config/blossom/router.yaml`。每个 route 指向一个 OpenAI-compatible endpoint。云端模型名不会编译进命令。API key 推荐通过 `api_key_env` 间接读取，不要直接写入 YAML。

每个 provider 也支持环境变量覆盖：

```bash
export BLOSSOM_LOCAL_MODEL=qwen3:8b
export BLOSSOM_LOCAL_BASE_URL=http://localhost:11434/v1
export BLOSSOM_STRONG_MODEL=your-model
export BLOSSOM_STRONG_BASE_URL=https://api.example.com/v1
export BLOSSOM_STRONG_API_KEY=secret
```

覆盖变量格式为 `BLOSSOM_<PROVIDER>_{MODEL,BASE_URL,API_KEY}`，环境变量优先于 YAML。

## 使用

```bash
bloom local "总结这段日志"
bloom ask "解释这次部署失败"
git diff | bloom code "审查这些修改"
bloom strong "比较这些架构方案"
bloom auto "翻译这份发布说明"
bloom auto --dry-run "设计一个迁移方案"
cat error.log | bloom auto
bloom serve
```

`auto` 使用可解释规则：短通用任务进入 local；摘要、翻译、分类和中等长度任务进入 cheap；编码关键词进入 normal；架构、安全、迁移、根因分析和超过 3,000 字符的提示词进入 strong。路由结果写入 stderr，`--dry-run` 绝不会调用模型。

`code` 是独立的可配置 provider slot。它可以连接 Codex/OpenAI-compatible chat-completions endpoint，但每次只进行一次请求，不做隐藏规划、重试或模型串联。

### 本地浏览器 UI

运行 `bloom serve` 打开 `http://127.0.0.1:7331`。本地控制台可以选择 route、粘贴 prompt、预览自动路由、执行请求并复制结果。它与 CLI 使用同一份 YAML 和环境变量；解析后的 API key 始终留在服务端。服务器会拒绝非 loopback 地址。

## 连接 Ollama 与 Qwen

安装并启动 Ollama，然后拉取配置中使用的模型：

```bash
brew install ollama
ollama serve
ollama pull qwen3:8b
bloom local "用三句话总结 Kubernetes readiness probe"
```

根据 Mac 内存选择 Qwen 大小，只需修改 `providers.local.model`。Ollama 的兼容 endpoint 默认为 `http://localhost:11434/v1`，不需要 API key。

## 云端 endpoint

为 `cheap`、`normal`、`strong` 和 `code` 设置各自的 `base_url`、`model` 与 `api_key_env`。不同层级可以使用同一厂商的不同模型，也可以使用不同的兼容厂商。项目不假定任何具体 OpenAI 型号。

## 范围

MVP 明确不包含 RAG、MCP、Web UI、数据库、对话历史、流式输出、工具调用和自动回退。只有真实使用证明必要时才进入第二阶段。

## 项目素材

- [公开项目站](https://alanthssss.github.io/blossom-router)
- [常见问题](FAQ.zh-CN.md)
- [实用性评估指南](docs/evaluation.zh-CN.md)
- [有效配置指南](docs/configuration-guide.zh-CN.md)
- [中文项目描述与传播文案](docs/project-description.zh-CN.md)
- [英文发布海报](docs/assets/blossom-router-poster-en.jpg)
- [中文发布海报](docs/assets/blossom-router-poster-zh-CN.jpg)
- [Agent 指南](AGENTS.zh-CN.md)与[交接记录](HANDOFF.zh-CN.md)
