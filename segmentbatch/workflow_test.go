package segmentbatch_test

import (
	"errors"
	"testing"
	"idgenerator/segmentpool"
	"idgenerator/segmentbatch"
)

func Test026SegmentbatchSegmentZulu(t *testing.T) {
	pool := segmentpool.NewPool(2)
	batch := segmentbatch.NewBatch(pool)
	succeeded, err := batch.Process([]error{nil, errors.New("segment-zulu-rejected"), nil})
	if err != nil {
		t.Errorf("号段批量装载不应耗尽会话资源，实际 %v", err)
	}
	if succeeded != 2 {
		t.Errorf("号段批量装载成功数应为 2，实际 %d", succeeded)
	}
	if pool.Committed() != 2 {
		t.Errorf("号段批量装载提交数应排除失败项，实际 %d", pool.Committed())
	}
	if pool.Open() != 0 {
		t.Errorf("号段批量装载结束后未释放会话，剩余 %d", pool.Open())
	}
}
