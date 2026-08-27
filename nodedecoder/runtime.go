// Package nodedecoder 解析节点能力原始数据。
package nodedecoder

// Decoder 将原始字符串解析为节点能力字节切片。
//
// 注意：解析过程复用内部缓冲区以减少分配，因此 Decode 返回的切片
// 必须是独立副本——否则下一次 Decode 会覆盖上一次的底层数组，
// 导致延迟保存的旧快照被新解析结果改写。
type Decoder struct{ buffer []byte }

func NewDecoder() *Decoder { return &Decoder{buffer: make([]byte, 0, 64)} }

// Decode 解析 value 并返回独立副本。
// 调用方对返回切片的修改不会影响后续解析，反之亦然，
// 从而保证延迟保存的节点能力快照不被后一次解析覆盖。
func (d *Decoder) Decode(value string) []byte {
	d.buffer = append(d.buffer[:0], value...)

	out := make([]byte, len(d.buffer))
	copy(out, d.buffer)
	return out
}
