package nodehealth_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"idgenerator/healthctx"
	"idgenerator/nodehealth"
)

func Test004NodehealthLifecycle(t *testing.T) {
	probe := healthctx.NewProbe(120 * time.Millisecond)
	runner := nodehealth.NewRunner(probe)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	started := time.Now()
	err := runner.Execute(ctx)
	elapsed := time.Since(started)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("节点健康确认取消后应返回 context.Canceled，实际 %v", err)
	}
	if elapsed > 60*time.Millisecond {
		t.Errorf("节点健康确认取消后仍等待下游，耗时 %s", elapsed)
	}
	if probe.Calls() != 1 {
		t.Errorf("节点健康确认取消后下游调用次数应为 1，实际 %d", probe.Calls())
	}
}
