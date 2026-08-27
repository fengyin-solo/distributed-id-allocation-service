package leaserenewal_test

import (
	"errors"
	"testing"

	"idgenerator/leasetransport"
	"idgenerator/leaserenewal"
)

func Test012LeaserenewalLeaseLima(t *testing.T) {
	gateway := leasetransport.NewGateway(leasetransport.ErrRejected, nil)
	coordinator := leaserenewal.NewCoordinator(gateway)
	err := coordinator.Run("lease-lima")
	if !errors.Is(err, leasetransport.ErrRejected) {
		t.Errorf("租约续期提交应保留拒绝错误，实际 %v", err)
	}
	if gateway.Calls() != 1 {
		t.Errorf("租约续期提交遇到拒绝后调用次数应为 1，实际 %d", gateway.Calls())
	}
	if coordinator.Committed() != 0 {
		t.Errorf("租约续期提交被拒后提交数应为 0，实际 %d", coordinator.Committed())
	}
}
