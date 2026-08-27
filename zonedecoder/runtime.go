package zonedecoder

type Decoder struct{ buffer []byte }

func NewDecoder() *Decoder { return &Decoder{buffer: make([]byte, 0, 64)} }

// Decode 解析心跳区域字符串。返回的字节切片与内部 buffer 互不共享底层内存，
// 多次调用互不影响：调用方可安全持有并使用前一次的返回值，即使后续再次调用 Decode。
func (d *Decoder) Decode(value string) []byte {
	d.buffer = append(d.buffer[:0], value...)
	out := make([]byte, len(d.buffer))
	copy(out, d.buffer)
	return out
}
