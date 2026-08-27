package segmentadmission_test

import (
	"testing"
	"idgenerator/segmentadmission"
)

func Test018SegmentadmissionSegmentRomeo(t *testing.T) {
	coordinator := segmentadmission.NewCoordinator(false)
	var applyErr error
	var panicValue any
	func() {
		defer func() { panicValue = recover() }()
		applyErr = coordinator.Apply("segment-romeo")
	}()
	if panicValue != nil {
		t.Errorf("号段准入规则缺省路径不应 panic，实际 %v", panicValue)
	}
	if applyErr != nil {
		t.Errorf("号段准入规则缺省路径应成功，实际 %v", applyErr)
	}
	if coordinator.Label("segment-romeo") != "active" {
		t.Errorf("号段准入规则缺省路径应保存 active 标签")
	}
}
