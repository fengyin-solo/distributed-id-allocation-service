package statspolicy

import "errors"

type Gate interface { Allow(string) error }
type ruleGate struct { blocked string }

func NewGate(enabled bool) Gate {
	if !enabled {
		// 附加规则缺省时返回真正的 nil 接口，避免 nil 指针被包装成非 nil 接口，
		// 否则调用方对接口的非 nil 判断失效，调用 Allow 会在 nil 接收者上 panic。
		return nil
	}
	return &ruleGate{blocked: "blocked"}
}

func (g *ruleGate) Allow(value string) error {
	if g.blocked == value { return errors.New("value blocked") }
	return nil
}
