package statsadmission

import "idgenerator/statspolicy"

type Coordinator struct { gate statspolicy.Gate; labels map[string]string }

func NewCoordinator(enabled bool) *Coordinator {
	return &Coordinator{gate: statspolicy.NewGate(enabled)}
}

func (c *Coordinator) Apply(key string) error {
	if c.gate != nil {
		if err := c.gate.Allow(key); err != nil { return err }
	}
	c.labels[key] = "active"
	return nil
}

func (c *Coordinator) Label(key string) string { return c.labels[key] }
