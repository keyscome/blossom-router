package router

import "strings"

type Decision struct{ Route, Reason string }

func Choose(prompt string) Decision {
	p := strings.ToLower(prompt)
	strong := []string{"architecture", "security review", "migration plan", "root cause", "架构", "安全审查", "迁移方案", "根因分析", "复杂推理"}
	code := []string{"implement", "refactor", "debug", "stack trace", "code review", "实现", "重构", "调试", "报错", "代码审查"}
	cheap := []string{"summarize", "translate", "classify", "format", "摘要", "翻译", "分类", "格式化"}
	if contains(p, strong) || len([]rune(prompt)) > 3000 {
		return Decision{"strong", "complexity keyword or very long prompt"}
	}
	if contains(p, code) {
		return Decision{"normal", "coding task"}
	}
	if contains(p, cheap) || len([]rune(prompt)) > 1200 {
		return Decision{"cheap", "batch-style or medium-length task"}
	}
	return Decision{"local", "short general task"}
}

func contains(s string, words []string) bool {
	for _, w := range words {
		if strings.Contains(s, w) {
			return true
		}
	}
	return false
}
