# Blossom Router 项目描述

## GitHub 简介

用一个轻量、可配置的 CLI，在本地与云端模型之间路由 AI 提示词。

## 简短介绍

Blossom Router 是面向 macOS 的轻量级 Go CLI，帮助开发者有意识地使用高价 AI 额度。通过 `bloom`，日常提示词可以交给本地 Ollama，普通和编程任务可以连接可配置的兼容 endpoint，困难任务再使用强模型；自动路由还可以在发送请求前预览决定。

## 完整介绍

并非每次 AI 工作都需要最强模型。Blossom Router 用一个小型命令行界面统一 local、cheap、normal、strong 和 code 五个 provider slot。所有云端模型、endpoint 和凭据均来自配置。自动路由使用公开的关键词与长度规则，不额外调用分类模型，因此路由本身不会消耗模型额度。

初始版本刻意保持克制：一个 Go 二进制、一个 YAML 配置、stdin 与命令行参数输入，以及 OpenAI-compatible chat-completions 协议。它没有隐藏重试、自动回退、遥测、对话数据库、RAG、MCP 或 Agent 框架。

## 标语

让每一次 AI 调用，都恰到好处。

## 建议 GitHub Topics

`ai-router`、`cli`、`golang`、`ollama`、`apple-silicon`、`local-llm`、`openai-compatible`、`developer-tools`

## 社交媒体文案

Blossom Router 来了：一个轻量的 `bloom` 命令，连接本地与云端 AI 模型。日常工作留在本地，把高价模型留给真正困难的任务，并在花费任何 token 前预览自动路由决定。
