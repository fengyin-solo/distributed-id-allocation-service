package nodeadmission_test

import (
	"testing"
	"idgenerator/nodeadmission"
)

func Test016NodeadmissionNodePapa(t *testing.T) {
	coordinator := nodeadmission.NewCoordinator(false)
	var applyErr error
	var panicValue any
	func() {
		defer func() { panicValue = recover() }()
		applyErr = coordinator.Apply("node-papa")
	}()
	if panicValue != nil {
		t.Errorf("节点接入策略缺省路径不应 panic，实际 %v", panicValue)
	}
	if applyErr != nil {
		t.Errorf("节点接入策略缺省路径应成功，实际 %v", applyErr)
	}
	if coordinator.Label("node-papa") != "active" {
		t.Errorf("节点接入策略缺省路径应保存 active 标签")
	}
}
