package allocationaudit_test

import (
	"errors"
	"testing"

	"idgenerator/audittransport"
	"idgenerator/allocationaudit"
)

func Test014AllocationauditAuditNovember(t *testing.T) {
	gateway := audittransport.NewGateway(audittransport.ErrRejected, nil)
	coordinator := allocationaudit.NewCoordinator(gateway)
	err := coordinator.Run("audit-november")
	if !errors.Is(err, audittransport.ErrRejected) {
		t.Errorf("分配审计提交应保留拒绝错误，实际 %v", err)
	}
	if gateway.Calls() != 1 {
		t.Errorf("分配审计提交遇到拒绝后调用次数应为 1，实际 %d", gateway.Calls())
	}
	if coordinator.Committed() != 0 {
		t.Errorf("分配审计提交被拒后提交数应为 0，实际 %d", coordinator.Committed())
	}
}
