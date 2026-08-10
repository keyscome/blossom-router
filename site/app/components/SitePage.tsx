"use client";

import { useMemo, useState } from "react";

type Locale = "en" | "zh";

const content = {
  en: {
    language: "中文",
    languageHref: "/zh",
    nav: ["How it works", "Evaluate", "Configure", "FAQ"],
    eyebrow: "BLOSSOM ECOSYSTEM · OPEN SOURCE",
    title: "One prompt.\nThe right model.",
    lead: "Keep routine AI work local. Route harder tasks to the cloud only when they deserve the cost.",
    primary: "Get started",
    secondary: "View on GitHub",
    command: "bloom auto --dry-run \"design a migration plan\"",
    routeLabel: "ROUTE DECISION",
    routeValue: "strong",
    reason: "complexity keyword",
    sectionHow: "A small router, not another AI platform.",
    sectionHowLead: "Blossom Router sits between your prompt and the models you already use. The rules are visible, local-first, and cheap to run.",
    routes: [
      ["Local", "Short questions and private routine work", "Ollama"],
      ["Cheap", "Summaries, translation, formatting, batch text", "Cloud · low cost"],
      ["Normal", "Everyday coding and general cloud work", "Cloud · balanced"],
      ["Strong", "Architecture, security, migration, root cause", "Cloud · premium"],
      ["Code", "An explicit slot for a coding-compatible endpoint", "One request only"],
    ],
    evalTitle: "Is Blossom Router useful for you?",
    evalLead: "Check what matches your workflow. The score is local to this page and is not stored.",
    checks: [
      "I use both local and cloud models",
      "I want to reduce premium-model quota usage",
      "I repeat summaries, translations, logs, or small code tasks",
      "I prefer explicit rules over a hidden AI classifier",
      "My providers expose OpenAI-compatible chat completions",
    ],
    scoreLabels: ["Optional", "Worth a trial", "Strong fit"],
    scoreNotes: ["A single provider may be simpler without routing.", "Start with local + normal and observe one week.", "This workflow matches the MVP design well."],
    configTitle: "Build a practical first configuration.",
    configLead: "Start with two useful tiers. Add strong and code only after real tasks justify them.",
    localModel: "Local Ollama model",
    cloudURL: "Cloud base URL",
    normalModel: "Normal cloud model",
    strongModel: "Strong cloud model (optional)",
    copy: "Copy YAML",
    copied: "Copied",
    configNote: "API keys stay in environment variables. This generator never asks for or stores a key.",
    faqTitle: "Frequently asked questions",
    faqs: [
      ["Does the public website call my models?", "No. The public site only explains and generates configuration text. Run bloom serve for the local operating UI."],
      ["Does automatic routing consume tokens?", "No. The MVP uses deterministic keyword and length rules. Dry-run never calls a provider."],
      ["Can it use Codex?", "The code slot accepts a configurable OpenAI-compatible chat-completions endpoint. It does not assume a proprietary transport."],
      ["Where are API keys stored?", "Use environment variables referenced by api_key_env. Keys are not compiled into bloom or exposed to the browser UI."],
      ["Which local model should I use?", "Choose a Qwen or other Ollama model that fits your Mac memory. Use a model you already run reliably before chasing benchmarks."],
      ["What should stay out of this MVP?", "RAG, MCP, databases, agent frameworks, hidden retries, and automatic fallback remain deferred until actual usage proves a need."],
    ],
    localTitle: "Prefer buttons to terminal commands?",
    localLead: "Run one command to open the local control panel. Choose a route, paste a task, preview the decision, and see the result without handling API keys in a webpage.",
    footer: "Blossom Router · lightweight, local-first, explainable",
  },
  zh: {
    language: "EN",
    languageHref: "/",
    nav: ["工作方式", "适用性评估", "配置", "常见问题"],
    eyebrow: "BLOSSOM 生态 · 开源项目",
    title: "一个提示词，\n交给合适的模型。",
    lead: "日常 AI 工作留在本地，真正困难的任务才值得调用云端强模型。",
    primary: "开始使用",
    secondary: "查看 GitHub",
    command: "bloom auto --dry-run \"设计一个迁移方案\"",
    routeLabel: "路由决定",
    routeValue: "strong",
    reason: "命中复杂任务关键词",
    sectionHow: "一个小型路由器，而不是另一个 AI 平台。",
    sectionHowLead: "Blossom Router 位于提示词和你已有模型之间。规则公开、本地优先，路由本身无需调用模型。",
    routes: [
      ["本地", "短问题、隐私内容和日常工作", "Ollama"],
      ["经济", "摘要、翻译、格式化和批量文本", "云端 · 低成本"],
      ["标准", "日常编程和普通云端任务", "云端 · 均衡"],
      ["强力", "架构、安全、迁移和根因分析", "云端 · 高能力"],
      ["编程", "单独配置兼容的编程 endpoint", "每次只请求一次"],
    ],
    evalTitle: "Blossom Router 适合你吗？",
    evalLead: "选择符合你工作方式的项目。评分只在当前页面计算，不会保存。",
    checks: [
      "我同时使用本地模型和云端模型",
      "我希望减少高价模型额度消耗",
      "我经常处理摘要、翻译、日志或小型代码任务",
      "我更喜欢明确规则，而不是隐藏的 AI 分类器",
      "我的 provider 支持 OpenAI-compatible chat completions",
    ],
    scoreLabels: ["暂非必需", "值得试用", "非常适合"],
    scoreNotes: ["如果只有一个 provider，直接调用可能更简单。", "先配置 local 和 normal，观察一周使用情况。", "你的工作方式与当前 MVP 非常匹配。"],
    configTitle: "生成一份实用的初始配置。",
    configLead: "先配置两个真正有用的层级。只有真实任务需要时，再增加 strong 和 code。",
    localModel: "本地 Ollama 模型",
    cloudURL: "云端 Base URL",
    normalModel: "标准云端模型",
    strongModel: "强力云端模型（可选）",
    copy: "复制 YAML",
    copied: "已复制",
    configNote: "API key 始终保存在环境变量中；配置生成器不会询问或保存密钥。",
    faqTitle: "常见问题",
    faqs: [
      ["公开网站会调用我的模型吗？", "不会。公开站只负责介绍和生成配置文本。真正操作模型请在本机运行 bloom serve。"],
      ["自动路由会消耗 token 吗？", "不会。MVP 使用确定性的关键词与长度规则，dry-run 绝不会调用 provider。"],
      ["可以接入 Codex 吗？", "code slot 可以连接可配置的 OpenAI-compatible chat-completions endpoint，不假定专有协议。"],
      ["API key 保存在哪里？", "使用 api_key_env 引用的环境变量。密钥不会编译进 bloom，也不会暴露给浏览器 UI。"],
      ["本地应该使用哪个模型？", "选择适合 Mac 内存、已经能够稳定运行的 Qwen 或其他 Ollama 模型，不必一开始就追逐跑分。"],
      ["MVP 暂时不应该加入什么？", "RAG、MCP、数据库、Agent 框架、隐藏重试和自动回退继续推迟，直到实际使用证明需要。"],
    ],
    localTitle: "不想记命令？使用本地操作页面。",
    localLead: "只需一个命令打开本地控制台。选择路由、粘贴任务、预览决定并查看结果，不需要在网页中处理 API key。",
    footer: "Blossom Router · 轻量、本地优先、可解释",
  },
};

export function SitePage({ locale }: { locale: Locale }) {
  const c = content[locale];
  const [checks, setChecks] = useState<boolean[]>(Array(5).fill(false));
  const [copied, setCopied] = useState(false);
  const [localModel, setLocalModel] = useState("qwen3:8b");
  const [cloudURL, setCloudURL] = useState("https://api.example.com/v1");
  const [normalModel, setNormalModel] = useState("your-normal-model");
  const [strongModel, setStrongModel] = useState("your-strong-model");
  const score = checks.filter(Boolean).length;
  const scoreBand = score < 2 ? 0 : score < 4 ? 1 : 2;
  const yaml = useMemo(() => `default_route: normal

providers:
  local:
    base_url: http://localhost:11434/v1
    model: ${localModel || "your-local-model"}
  normal:
    base_url: ${cloudURL || "https://api.example.com/v1"}
    model: ${normalModel || "your-normal-model"}
    api_key_env: BLOSSOM_NORMAL_API_KEY
  strong:
    base_url: ${cloudURL || "https://api.example.com/v1"}
    model: ${strongModel || "your-strong-model"}
    api_key_env: BLOSSOM_STRONG_API_KEY`, [localModel, cloudURL, normalModel, strongModel]);

  async function copyYaml() {
    await navigator.clipboard.writeText(yaml);
    setCopied(true);
    setTimeout(() => setCopied(false), 1400);
  }

  return <main>
    <nav className="nav shell">
      <a className="brand" href={locale === "en" ? "/" : "/zh"}><span>✦</span> Blossom Router</a>
      <div className="navlinks">
        <a href="#how">{c.nav[0]}</a><a href="#evaluate">{c.nav[1]}</a><a href="#configure">{c.nav[2]}</a><a href="#faq">{c.nav[3]}</a>
      </div>
      <a className="language" href={c.languageHref}>{c.language}</a>
    </nav>

    <section className="hero shell">
      <div className="hero-copy"><p className="eyebrow">{c.eyebrow}</p><h1>{c.title.split("\n").map((line, i) => <span key={i}>{line}</span>)}</h1><p className="lead">{c.lead}</p><div className="hero-actions"><a className="button primary" href="#configure">{c.primary}</a><a className="button secondary" href="https://github.com/keyscome/blossom-router">{c.secondary}</a></div></div>
      <div className="route-demo"><div className="flower" aria-hidden="true"><i></i><i></i><i></i><i></i><i></i><b>✦</b></div><div className="terminal"><div className="terminal-top"><span></span><span></span><span></span></div><code>$ {c.command}</code><div className="decision"><small>{c.routeLabel}</small><strong>{c.routeValue}</strong><span>{c.reason}</span></div></div></div>
    </section>

    <section id="how" className="section shell"><div className="section-heading"><p className="eyebrow">01 · ROUTES</p><h2>{c.sectionHow}</h2><p>{c.sectionHowLead}</p></div><div className="route-grid">{c.routes.map((route, i) => <article key={route[0]}><span className={`route-dot r${i}`}></span><h3>{route[0]}</h3><p>{route[1]}</p><small>{route[2]}</small></article>)}</div></section>

    <section id="evaluate" className="section shell split"><div className="section-heading"><p className="eyebrow">02 · FIT CHECK</p><h2>{c.evalTitle}</h2><p>{c.evalLead}</p></div><div className="score-card"><div className="score-ring" style={{"--score": `${score * 20}%`} as React.CSSProperties}><strong>{score}/5</strong></div><div><h3>{c.scoreLabels[scoreBand]}</h3><p>{c.scoreNotes[scoreBand]}</p></div></div><div className="checklist">{c.checks.map((label, i) => <label key={label}><input type="checkbox" checked={checks[i]} onChange={() => setChecks(v => v.map((x, j) => j === i ? !x : x))}/><span>{label}</span></label>)}</div></section>

    <section id="configure" className="section shell"><div className="section-heading"><p className="eyebrow">03 · CONFIGURE</p><h2>{c.configTitle}</h2><p>{c.configLead}</p></div><div className="config-grid"><div className="config-form"><label>{c.localModel}<input value={localModel} onChange={e => setLocalModel(e.target.value)}/></label><label>{c.cloudURL}<input value={cloudURL} onChange={e => setCloudURL(e.target.value)}/></label><label>{c.normalModel}<input value={normalModel} onChange={e => setNormalModel(e.target.value)}/></label><label>{c.strongModel}<input value={strongModel} onChange={e => setStrongModel(e.target.value)}/></label><p>{c.configNote}</p></div><div className="code-card"><div className="code-head"><span>router.yaml</span><button onClick={copyYaml}>{copied ? c.copied : c.copy}</button></div><pre>{yaml}</pre></div></div></section>

    <section className="local-ui shell"><div><p className="eyebrow">LOCAL UI</p><h2>{c.localTitle}</h2><p>{c.localLead}</p></div><code>bloom serve</code></section>

    <section id="faq" className="section shell"><div className="section-heading"><p className="eyebrow">04 · FAQ</p><h2>{c.faqTitle}</h2></div><div className="faq-list">{c.faqs.map(([q, a]) => <details key={q}><summary>{q}</summary><p>{a}</p></details>)}</div></section>
    <footer className="shell"><span>{c.footer}</span><a href="https://github.com/keyscome/blossom-router">GitHub ↗</a></footer>
  </main>;
}
