package allocationbatch

import "idgenerator/writerpool"

type Batch struct { pool *writerpool.Pool }

func NewBatch(pool *writerpool.Pool) *Batch { return &Batch{pool: pool} }

func (b *Batch) Process(outcomes []error) (int, error) {
	succeeded := 0
	for _, outcome := range outcomes {
		session, err := b.pool.Acquire()
		if err != nil { return succeeded, err }
		success := outcome == nil
		if success { succeeded++ }
		// Close promptly so the session is returned to the pool within this
		// iteration; deferring would let open sessions accumulate until the
		// whole batch returns, exhausting capacity mid-batch. A rejected
		// record (success == false) is still closed, but not committed.
		session.Close(success)
	}
	return succeeded, nil
}
