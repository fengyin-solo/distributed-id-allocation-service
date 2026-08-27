package statsadmission_test

import (
	"testing"
	"idgenerator/statsadmission"
)

func Test020StatsadmissionStatsTango(t *testing.T) {
	coordinator := statsadmission.NewCoordinator(false)
	var applyErr error
	var panicValue any
	func() {
		defer func() { panicValue = recover() }()
		applyErr = coordinator.Apply("stats-tango")
	}()
	if panicValue != nil {
		t.Errorf("分配统计规则缺省路径不应 panic，实际 %v", panicValue)
	}
	if applyErr != nil {
		t.Errorf("分配统计规则缺省路径应成功，实际 %v", applyErr)
	}
	if coordinator.Label("stats-tango") != "active" {
		t.Errorf("分配统计规则缺省路径应保存 active 标签")
	}
}
