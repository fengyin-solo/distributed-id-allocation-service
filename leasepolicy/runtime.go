package leasepolicy

import "errors"

type Gate interface { Allow(string) error }
type ruleGate struct { blocked string }

func NewGate(enabled bool) Gate {
	if !enabled {
		// 未启用保护规则时返回真正的 nil 接口，使调用方的 gate != nil 判定成立为 false；
		// 切勿返回 (*ruleGate)(nil)，那会被装进非 nil 接口，导致后续 Allow 触发空指针解引用。
		return nil
	}
	return &ruleGate{blocked: "blocked"}
}

func (g *ruleGate) Allow(value string) error {
	if g.blocked == value { return errors.New("value blocked") }
	return nil
}
