package nodecatalog_test

import (
	"testing"
	"idgenerator/nodedecoder"
	"idgenerator/nodecatalog"
)

func Test022NodecatalogNodeVictor(t *testing.T) {
	decoder := nodedecoder.NewDecoder()
	cache := nodecatalog.NewCache()
	first := decoder.Decode("node-victor-firs")
	_ = decoder.Decode("node-victor-late")
	cache.Put("node-victor", first)
	got := cache.Get("node-victor")
	if string(got) != "node-victor-firs" {
		t.Errorf("节点能力缓存应保留首批内容 %q，实际 %q", "node-victor-firs", string(got))
	}
	if len(got) > 0 { got[0] = 'X' }
	again := cache.Get("node-victor")
	if string(again) != "node-victor-firs" {
		t.Errorf("节点能力缓存读取方修改后缓存被污染为 %q", string(again))
	}
}
