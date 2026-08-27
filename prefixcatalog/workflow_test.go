package prefixcatalog_test

import (
	"testing"
	"idgenerator/prefixdecoder"
	"idgenerator/prefixcatalog"
)

func Test024PrefixcatalogPrefixXray(t *testing.T) {
	decoder := prefixdecoder.NewDecoder()
	cache := prefixcatalog.NewCache()
	first := decoder.Decode("prefix-xray-firs")
	_ = decoder.Decode("prefix-xray-late")
	cache.Put("prefix-xray", first)
	got := cache.Get("prefix-xray")
	if string(got) != "prefix-xray-firs" {
		t.Errorf("号码前缀缓存应保留首批内容 %q，实际 %q", "prefix-xray-firs", string(got))
	}
	if len(got) > 0 { got[0] = 'X' }
	again := cache.Get("prefix-xray")
	if string(again) != "prefix-xray-firs" {
		t.Errorf("号码前缀缓存读取方修改后缓存被污染为 %q", string(again))
	}
}
