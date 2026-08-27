package noderegister_test

import (
	"errors"
	"testing"

	"idgenerator/nodetransport"
	"idgenerator/noderegister"
)

func Test011NoderegisterNodeKilo(t *testing.T) {
	gateway := nodetransport.NewGateway(nodetransport.ErrRejected, nil)
	coordinator := noderegister.NewCoordinator(gateway)
	err := coordinator.Run("node-kilo")
	if !errors.Is(err, nodetransport.ErrRejected) {
		t.Errorf("节点注册应保留拒绝错误，实际 %v", err)
	}
	if gateway.Calls() != 1 {
		t.Errorf("节点注册遇到拒绝后调用次数应为 1，实际 %d", gateway.Calls())
	}
	if coordinator.Committed() != 0 {
		t.Errorf("节点注册被拒后提交数应为 0，实际 %d", coordinator.Committed())
	}
}
