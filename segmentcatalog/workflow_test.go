package segmentcatalog_test

import (
	"testing"
	"idgenerator/segmentdecoder"
	"idgenerator/segmentcatalog"
)

func Test021SegmentcatalogSegmentUniform(t *testing.T) {
	decoder := segmentdecoder.NewDecoder()
	cache := segmentcatalog.NewCache()
	first := decoder.Decode("segment-uniform-")
	_ = decoder.Decode("segment-uniform-")
	cache.Put("segment-uniform", first)
	got := cache.Get("segment-uniform")
	if string(got) != "segment-uniform-" {
		t.Errorf("号段标签缓存应保留首批内容 %q，实际 %q", "segment-uniform-", string(got))
	}
	if len(got) > 0 { got[0] = 'X' }
	again := cache.Get("segment-uniform")
	if string(again) != "segment-uniform-" {
		t.Errorf("号段标签缓存读取方修改后缓存被污染为 %q", string(again))
	}
}
