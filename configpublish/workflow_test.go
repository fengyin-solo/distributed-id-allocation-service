package configpublish_test

import (
	"errors"
	"testing"

	"idgenerator/configtransport"
	"idgenerator/configpublish"
)

func Test015ConfigpublishConfigOscar(t *testing.T) {
	gateway := configtransport.NewGateway(configtransport.ErrRejected, nil)
	coordinator := configpublish.NewCoordinator(gateway)
	err := coordinator.Run("config-oscar")
	if !errors.Is(err, configtransport.ErrRejected) {
		t.Errorf("雪花配置发布应保留拒绝错误，实际 %v", err)
	}
	if gateway.Calls() != 1 {
		t.Errorf("雪花配置发布遇到拒绝后调用次数应为 1，实际 %d", gateway.Calls())
	}
	if coordinator.Committed() != 0 {
		t.Errorf("雪花配置发布被拒后提交数应为 0，实际 %d", coordinator.Committed())
	}
}
