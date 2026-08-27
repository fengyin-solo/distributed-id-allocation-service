package heartbeatfanin_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"idgenerator/heartbeatstream"
	"idgenerator/heartbeatfanin"
)

func Test009HeartbeatfaninHeartbeatIndia(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Millisecond)
	defer cancel()
	type result struct { items []string; err error }
	done := make(chan result, 1)
	started := time.Now()
	go func() {
		items, err := heartbeatfanin.Collect(ctx, []string{"heartbeat-india-one", "heartbeat-india-two", "heartbeat-india-three"}, 1)
		done <- result{items: items, err: err}
	}()
	select {
	case got := <-done:
		if !errors.Is(got.err, heartbeatstream.ErrSource) {
			t.Errorf("心跳流汇集应返回源错误，实际 %v", got.err)
		}
		if len(got.items) != 1 {
			t.Errorf("心跳流汇集应保留一条已收结果，实际 %d", len(got.items))
		}
		if time.Since(started) > 80*time.Millisecond {
			t.Errorf("心跳流汇集在源失败后返回过慢")
		}
	case <-time.After(240 * time.Millisecond):
		t.Fatalf("心跳流汇集来源失败后没有在超时边界内结束")
	}
}
