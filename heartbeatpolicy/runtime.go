package heartbeatpolicy

import "errors"

type Gate interface { Allow(string) error }
type ruleGate struct { blocked string }

func NewGate(enabled bool) Gate {
	if !enabled {
		// 未配置补充校验器时返回 nil 接口，避免 typed-nil 指针被赋给
		// Gate 接口后变成非 nil，从而误调 Allow 触发空指针 panic。
		return nil
	}
	return &ruleGate{blocked: "blocked"}
}

func (g *ruleGate) Allow(value string) error {
	if g.blocked == value { return errors.New("value blocked") }
	return nil
}
