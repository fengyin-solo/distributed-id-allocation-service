package heartbeatcheck_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"idgenerator/heartbeatctx"
	"idgenerator/heartbeatcheck"
)

func Test002HeartbeatcheckLifecycle(t *testing.T) {
	probe := heartbeatctx.NewProbe(120 * time.Millisecond)
	runner := heartbeatcheck.NewRunner(probe)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	started := time.Now()
	err := runner.Execute(ctx)
	elapsed := time.Since(started)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("心跳检查取消后应返回 context.Canceled，实际 %v", err)
	}
	if elapsed > 60*time.Millisecond {
		t.Errorf("心跳检查取消后仍等待下游，耗时 %s", elapsed)
	}
	if probe.Calls() != 1 {
		t.Errorf("心跳检查取消后下游调用次数应为 1，实际 %d", probe.Calls())
	}
}
