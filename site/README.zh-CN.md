# Blossom Router 站点

[English](README.md) | [简体中文](README.zh-CN.md)

本目录包含 Blossom Router 的中英双语公开项目站。英文页面位于 `/`，简体中文页面位于 `/zh`。

公开站点介绍路由模型，提供只在浏览器本地计算的适用性评估与 YAML 配置生成器，回答常见问题，并引导用户使用 localhost 操作 UI。它绝不接收 API key，也不会调用已配置的 AI provider。

## 开发与验证

使用 Node.js 22.13 或更高版本：

```bash
npm install
npm run dev
npm test
npm run lint
```

项目使用 Vinext 和 Sites 构建契约。托管项目标识保存在 `.openai/hosting.json`；不得在其中保存密钥或运行时凭据。
