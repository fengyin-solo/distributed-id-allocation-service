package configbatch

import "idgenerator/loaderpool"

type Batch struct { pool *loaderpool.Pool }

func NewBatch(pool *loaderpool.Pool) *Batch { return &Batch{pool: pool} }

func (b *Batch) Process(outcomes []error) (int, error) {
	succeeded := 0
	for _, outcome := range outcomes {
		session, err := b.pool.Acquire()
		if err != nil { return succeeded, err }
		success := outcome == nil
		session.Close(success) // release this config's session now, not at function return
		if success { succeeded++ }
	}
	return succeeded, nil
}
