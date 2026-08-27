package leaseadmission

import "idgenerator/leasepolicy"

type Coordinator struct { gate leasepolicy.Gate; labels map[string]string }

func NewCoordinator(enabled bool) *Coordinator {
	return &Coordinator{gate: leasepolicy.NewGate(enabled)}
}

func (c *Coordinator) Apply(key string) error {
	if c.gate != nil {
		if err := c.gate.Allow(key); err != nil { return err }
	}
	c.labels[key] = "active"
	return nil
}

func (c *Coordinator) Label(key string) string { return c.labels[key] }
