package leaseadmission_test

import (
	"testing"
	"idgenerator/leaseadmission"
)

func Test017LeaseadmissionLeaseQuebec(t *testing.T) {
	coordinator := leaseadmission.NewCoordinator(false)
	var applyErr error
	var panicValue any
	func() {
		defer func() { panicValue = recover() }()
		applyErr = coordinator.Apply("lease-quebec")
	}()
	if panicValue != nil {
		t.Errorf("租约保护策略缺省路径不应 panic，实际 %v", panicValue)
	}
	if applyErr != nil {
		t.Errorf("租约保护策略缺省路径应成功，实际 %v", applyErr)
	}
	if coordinator.Label("lease-quebec") != "active" {
		t.Errorf("租约保护策略缺省路径应保存 active 标签")
	}
}
