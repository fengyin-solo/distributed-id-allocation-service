package segmentadmission

import "idgenerator/segmentpolicy"

type Coordinator struct {
	gate   segmentpolicy.Gate
	labels map[string]string
}

func NewCoordinator(enabled bool) *Coordinator {
	return &Coordinator{
		gate:   segmentpolicy.NewGate(enabled),
		labels: make(map[string]string),
	}
}

// Apply 将号段 key 交给准入闸门校验，通过后写入使用标记 "active"。
// 当号段没有自定义准入规则（缺省）时，gate 为 nil，直接跳过校验并写入标记，
// 使号段正常进入使用状态。校验失败则不写入标记。
func (c *Coordinator) Apply(key string) error {
	if c.gate != nil {
		if err := c.gate.Allow(key); err != nil {
			return err
		}
	}
	c.labels[key] = "active"
	return nil
}

func (c *Coordinator) Label(key string) string { return c.labels[key] }
