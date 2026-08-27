package probebatch_test

import (
	"errors"
	"testing"
	"idgenerator/probepool"
	"idgenerator/probebatch"
)

func Test028ProbebatchProbeBirch(t *testing.T) {
	pool := probepool.NewPool(2)
	batch := probebatch.NewBatch(pool)
	succeeded, err := batch.Process([]error{nil, errors.New("probe-birch-rejected"), nil})
	if err != nil {
		t.Errorf("心跳批量探测不应耗尽会话资源，实际 %v", err)
	}
	if succeeded != 2 {
		t.Errorf("心跳批量探测成功数应为 2，实际 %d", succeeded)
	}
	if pool.Committed() != 2 {
		t.Errorf("心跳批量探测提交数应排除失败项，实际 %d", pool.Committed())
	}
	if pool.Open() != 0 {
		t.Errorf("心跳批量探测结束后未释放会话，剩余 %d", pool.Open())
	}
}
