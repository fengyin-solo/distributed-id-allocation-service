package statsadmission

import "idgenerator/statspolicy"

type Coordinator struct { gate statspolicy.Gate; labels map[string]string }

func NewCoordinator(enabled bool) *Coordinator {
	return &Coordinator{gate: statspolicy.NewGate(enabled), labels: map[string]string{}}
}

func (c *Coordinator) Apply(key string) error {
	if c.gate != nil {
		if err := c.gate.Allow(key); err != nil { return err }
	}
	// 附加规则缺省（gate 为 nil）时仍需写入当天状态标签，不能跳过。
	c.labels[key] = "active"
	return nil
}

func (c *Coordinator) Label(key string) string { return c.labels[key] }
