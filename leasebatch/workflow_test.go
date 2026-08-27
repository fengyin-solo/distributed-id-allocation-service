package leasebatch_test

import (
	"errors"
	"testing"
	"idgenerator/leasepool"
	"idgenerator/leasebatch"
)

func Test027LeasebatchLeaseAmber(t *testing.T) {
	pool := leasepool.NewPool(2)
	batch := leasebatch.NewBatch(pool)
	succeeded, err := batch.Process([]error{nil, errors.New("lease-amber-rejected"), nil})
	if err != nil {
		t.Errorf("租约批量续签不应耗尽会话资源，实际 %v", err)
	}
	if succeeded != 2 {
		t.Errorf("租约批量续签成功数应为 2，实际 %d", succeeded)
	}
	if pool.Committed() != 2 {
		t.Errorf("租约批量续签提交数应排除失败项，实际 %d", pool.Committed())
	}
	if pool.Open() != 0 {
		t.Errorf("租约批量续签结束后未释放会话，剩余 %d", pool.Open())
	}
}
