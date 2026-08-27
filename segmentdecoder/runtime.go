package segmentdecoder

// Decoder parses segment tag strings into byte slices.
type Decoder struct{}

// NewDecoder returns a stateless Decoder.
//
// A previous implementation reused a single internal buffer and returned
// slices that aliased it: decoding a second batch of tags overwrote the byte
// slices a caller was still holding from the first batch. Decode therefore
// returns a fresh, independent allocation each call — safe to retain and
// cache without leaking across batches.
func NewDecoder() *Decoder { return &Decoder{} }

// Decode returns an independent copy of the parsed tag bytes.
func (d *Decoder) Decode(value string) []byte {
	out := make([]byte, len(value))
	copy(out, value)
	return out
}
