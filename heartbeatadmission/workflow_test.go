package heartbeatadmission_test

import (
	"testing"
	"idgenerator/heartbeatadmission"
)

func Test019HeartbeatadmissionHeartbeatSierra(t *testing.T) {
	coordinator := heartbeatadmission.NewCoordinator(false)
	var applyErr error
	var panicValue any
	func() {
		defer func() { panicValue = recover() }()
		applyErr = coordinator.Apply("heartbeat-sierra")
	}()
	if panicValue != nil {
		t.Errorf("心跳补充规则缺省路径不应 panic，实际 %v", panicValue)
	}
	if applyErr != nil {
		t.Errorf("心跳补充规则缺省路径应成功，实际 %v", applyErr)
	}
	if coordinator.Label("heartbeat-sierra") != "active" {
		t.Errorf("心跳补充规则缺省路径应保存 active 标签")
	}
}
