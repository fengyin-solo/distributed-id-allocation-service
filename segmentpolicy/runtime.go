package segmentpolicy

import "errors"

type Gate interface { Allow(string) error }
type ruleGate struct { blocked string }

// NewGate 当准入未启用（无自定义准入规则）时返回 nil，使调用方通过 nil 判定跳过校验，
// 缺省规则下号段可直接进入使用状态。注意返回的是无类型的 nil 接口，而非带类型的 nil 指针，
// 否则 Coordinator 中 c.gate != nil 的判断会因 typed-nil 陷阱失效而触发空指针。
func NewGate(enabled bool) Gate {
	if !enabled {
		return nil
	}
	return &ruleGate{blocked: "blocked"}
}

func (g *ruleGate) Allow(value string) error {
	if g.blocked == value { return errors.New("value blocked") }
	return nil
}
