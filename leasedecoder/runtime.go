package leasedecoder

type Decoder struct { buffer []byte }

func NewDecoder() *Decoder { return &Decoder{buffer: make([]byte, 0, 64)} }

// Decode 解析租约范围并返回独立副本。
// 复用内部 buffer 只是为了避免每次分配，但返回值必须拷贝，
// 否则连续解析两组等长范围时首组结果会被后组覆盖。
func (d *Decoder) Decode(value string) []byte {
	d.buffer = append(d.buffer[:0], value...)
	out := make([]byte, len(d.buffer))
	copy(out, d.buffer)
	return out
}
