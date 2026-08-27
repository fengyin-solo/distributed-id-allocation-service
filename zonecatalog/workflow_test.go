package zonecatalog_test

import (
	"testing"
	"idgenerator/zonedecoder"
	"idgenerator/zonecatalog"
)

func Test025ZonecatalogZoneYankee(t *testing.T) {
	decoder := zonedecoder.NewDecoder()
	cache := zonecatalog.NewCache()
	first := decoder.Decode("zone-yankee-firs")
	_ = decoder.Decode("zone-yankee-late")
	cache.Put("zone-yankee", first)
	got := cache.Get("zone-yankee")
	if string(got) != "zone-yankee-firs" {
		t.Errorf("心跳区域缓存应保留首批内容 %q，实际 %q", "zone-yankee-firs", string(got))
	}
	if len(got) > 0 { got[0] = 'X' }
	again := cache.Get("zone-yankee")
	if string(again) != "zone-yankee-firs" {
		t.Errorf("心跳区域缓存读取方修改后缓存被污染为 %q", string(again))
	}
}
