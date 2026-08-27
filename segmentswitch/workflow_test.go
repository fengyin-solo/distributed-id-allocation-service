package segmentswitch_test

import (
	"errors"
	"testing"

	"idgenerator/segmenttransport"
	"idgenerator/segmentswitch"
)

func Test013SegmentswitchSegmentMike(t *testing.T) {
	gateway := segmenttransport.NewGateway(segmenttransport.ErrRejected, nil)
	coordinator := segmentswitch.NewCoordinator(gateway)
	err := coordinator.Run("segment-mike")
	if !errors.Is(err, segmenttransport.ErrRejected) {
		t.Errorf("号段切换应保留拒绝错误，实际 %v", err)
	}
	if gateway.Calls() != 1 {
		t.Errorf("号段切换遇到拒绝后调用次数应为 1，实际 %d", gateway.Calls())
	}
	if coordinator.Committed() != 0 {
		t.Errorf("号段切换被拒后提交数应为 0，实际 %d", coordinator.Committed())
	}
}
