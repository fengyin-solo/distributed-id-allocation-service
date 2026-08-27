package allocationbatch_test

import (
	"errors"
	"testing"
	"idgenerator/writerpool"
	"idgenerator/allocationbatch"
)

func Test029AllocationbatchAllocCedar(t *testing.T) {
	pool := writerpool.NewPool(2)
	batch := allocationbatch.NewBatch(pool)
	succeeded, err := batch.Process([]error{nil, errors.New("alloc-cedar-rejected"), nil})
	if err != nil {
		t.Errorf("分配记录批写不应耗尽会话资源，实际 %v", err)
	}
	if succeeded != 2 {
		t.Errorf("分配记录批写成功数应为 2，实际 %d", succeeded)
	}
	if pool.Committed() != 2 {
		t.Errorf("分配记录批写提交数应排除失败项，实际 %d", pool.Committed())
	}
	if pool.Open() != 0 {
		t.Errorf("分配记录批写结束后未释放会话，剩余 %d", pool.Open())
	}
}
