package leasebatch

import "idgenerator/leasepool"

type Batch struct { pool *leasepool.Pool }

func NewBatch(pool *leasepool.Pool) *Batch { return &Batch{pool: pool} }

func (b *Batch) Process(outcomes []error) (int, error) {
	succeeded := 0
	for _, outcome := range outcomes {
		session, err := b.pool.Acquire()
		if err != nil { return succeeded, err }
		success := outcome == nil
		session.Close(success)
		if success { succeeded++ }
	}
	return succeeded, nil
}
