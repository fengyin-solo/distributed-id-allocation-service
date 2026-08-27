package configbatch_test

import (
	"errors"
	"testing"
	"idgenerator/loaderpool"
	"idgenerator/configbatch"
)

func Test030ConfigbatchConfigDawn(t *testing.T) {
	pool := loaderpool.NewPool(2)
	batch := configbatch.NewBatch(pool)
	succeeded, err := batch.Process([]error{nil, errors.New("config-dawn-rejected"), nil})
	if err != nil {
		t.Errorf("配置批量加载不应耗尽会话资源，实际 %v", err)
	}
	if succeeded != 2 {
		t.Errorf("配置批量加载成功数应为 2，实际 %d", succeeded)
	}
	if pool.Committed() != 2 {
		t.Errorf("配置批量加载提交数应排除失败项，实际 %d", pool.Committed())
	}
	if pool.Open() != 0 {
		t.Errorf("配置批量加载结束后未释放会话，剩余 %d", pool.Open())
	}
}
