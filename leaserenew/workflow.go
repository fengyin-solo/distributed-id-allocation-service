package leaserenew

import (
	"context"
	"idgenerator/leasectx"
)

type Runner struct { probe *leasectx.Probe }

func NewRunner(probe *leasectx.Probe) *Runner { return &Runner{probe: probe} }

func (r *Runner) Execute(ctx context.Context) error {
	var last error
	for attempt := 0; attempt < 3; attempt++ {
		err := r.probe.Fetch(context.Background())
		if err == nil {
			return nil
		}
		last = err
	}
	return last
}
