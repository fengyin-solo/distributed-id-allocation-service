package allocationcollect_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"idgenerator/allocationstream"
	"idgenerator/allocationcollect"
)

func Test010AllocationcollectAllocJuliet(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Millisecond)
	defer cancel()
	type result struct { items []string; err error }
	done := make(chan result, 1)
	started := time.Now()
	go func() {
		items, err := allocationcollect.Collect(ctx, []string{"alloc-juliet-one", "alloc-juliet-two", "alloc-juliet-three"}, 1)
		done <- result{items: items, err: err}
	}()
	select {
	case got := <-done:
		if !errors.Is(got.err, allocationstream.ErrSource) {
			t.Errorf("分配回执汇集应返回源错误，实际 %v", got.err)
		}
		if len(got.items) != 1 {
			t.Errorf("分配回执汇集应保留一条已收结果，实际 %d", len(got.items))
		}
		if time.Since(started) > 80*time.Millisecond {
			t.Errorf("分配回执汇集在源失败后返回过慢")
		}
	case <-time.After(240 * time.Millisecond):
		t.Fatalf("分配回执汇集来源失败后没有在超时边界内结束")
	}
}
