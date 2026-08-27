package heartbeatadmission

import "idgenerator/heartbeatpolicy"

type Coordinator struct { gate heartbeatpolicy.Gate; labels map[string]string }

func NewCoordinator(enabled bool) *Coordinator {
	return &Coordinator{
		gate:   heartbeatpolicy.NewGate(enabled),
		labels: make(map[string]string),
	}
}

func (c *Coordinator) Apply(key string) error {
	if c.gate != nil {
		if err := c.gate.Allow(key); err != nil { return err }
	}
	c.labels[key] = "active"
	return nil
}

func (c *Coordinator) Label(key string) string { return c.labels[key] }
