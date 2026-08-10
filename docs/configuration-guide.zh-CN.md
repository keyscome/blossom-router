# 有效配置指南

[English](configuration-guide.md) | [简体中文](configuration-guide.zh-CN.md)

## 推荐的第一份配置

使用四个自动路由 slot，加一个可选的明确编程 slot：

- `local`：适合内存且运行稳定的 Ollama 模型。
- `cheap`：翻译、摘要等低成本云端任务。
- `normal`：均衡的通用与编程任务。
- `strong`：困难推理、架构、安全与根因分析。
- `code`：可选的明确编程 endpoint；`auto` 不会选择它。

如果暂时没有四个不同模型，可以让 cheap 与 normal 指向同一 endpoint。完整性比刻意区分更重要，因为 `auto` 不会静默回退，而会在缺少所选 route 时返回清晰错误。

## 示例

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

密钥保存在 YAML 之外：

```bash
export BLOSSOM_CHEAP_API_KEY=...
export BLOSSOM_NORMAL_API_KEY=...
export BLOSSOM_STRONG_API_KEY=...
export BLOSSOM_CODE_API_KEY=...
```

运行 `bloom serve` 启动低复杂度操作界面。它读取同一份配置，且不会把解析后的 key 返回浏览器。
