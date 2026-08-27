package leasecatalog_test

import (
	"testing"
	"idgenerator/leasedecoder"
	"idgenerator/leasecatalog"
)

func Test023LeasecatalogLeaseWhiskey(t *testing.T) {
	decoder := leasedecoder.NewDecoder()
	cache := leasecatalog.NewCache()
	first := decoder.Decode("lease-whiskey-fi")
	_ = decoder.Decode("lease-whiskey-la")
	cache.Put("lease-whiskey", first)
	got := cache.Get("lease-whiskey")
	if string(got) != "lease-whiskey-fi" {
		t.Errorf("租约范围缓存应保留首批内容 %q，实际 %q", "lease-whiskey-fi", string(got))
	}
	if len(got) > 0 { got[0] = 'X' }
	again := cache.Get("lease-whiskey")
	if string(again) != "lease-whiskey-fi" {
		t.Errorf("租约范围缓存读取方修改后缓存被污染为 %q", string(again))
	}
}
