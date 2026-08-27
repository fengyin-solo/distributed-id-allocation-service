package prefixdecoder

type Decoder struct { buffer []byte }

func NewDecoder() *Decoder { return &Decoder{buffer: make([]byte, 0, 64)} }

// Decode materializes value into the decoder's reusable scratch buffer and
// returns an independent copy. The returned slice does not alias the internal
// buffer, so subsequent Decode calls cannot mutate previously returned results
// and each prefix keeps its own content.
func (d *Decoder) Decode(value string) []byte {
	d.buffer = append(d.buffer[:0], value...)
	out := make([]byte, len(d.buffer))
	copy(out, d.buffer)
	return out
}
