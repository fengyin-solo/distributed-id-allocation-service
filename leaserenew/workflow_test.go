package leaserenew_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"idgenerator/leasectx"
	"idgenerator/leaserenew"
)

func Test001LeaserenewLifecycle(t *testing.T) {
	probe := leasectx.NewProbe(120 * time.Millisecond)
	runner := leaserenew.NewRunner(probe)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	started := time.Now()
	err := runner.Execute(ctx)
	elapsed := time.Since(started)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("租约续期取消后应返回 context.Canceled，实际 %v", err)
	}
	if elapsed > 60*time.Millisecond {
		t.Errorf("租约续期取消后仍等待下游，耗时 %s", elapsed)
	}
	if probe.Calls() != 1 {
		t.Errorf("租约续期取消后下游调用次数应为 1，实际 %d", probe.Calls())
	}
}
