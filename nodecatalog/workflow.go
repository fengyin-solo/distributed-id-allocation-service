// Package nodecatalog 缓存已解析的节点能力快照。
package nodecatalog

import "sync"

// Cache 以 key-value 形式缓存节点能力快照。
// 缓存中的值是不可变的快照：写入时拷贝，读取时也返回独立副本，
// 避免外部或后续请求对底层数组的修改写回到缓存，从而覆盖已有快照。
type Cache struct {
	mu     sync.RWMutex
	values map[string][]byte
}

func NewCache() *Cache { return &Cache{values: make(map[string][]byte)} }

// Put 保存节点能力快照。传入的 value 会被拷贝一份后存入缓存，
// 因此调用方之后修改原切片不会影响缓存内容。
func (c *Cache) Put(key string, value []byte) {
	snap := make([]byte, len(value))
	copy(snap, value)

	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[key] = snap
}

// Get 返回 key 对应快照的独立副本。返回的切片与缓存内部数据互不影响，
// 调用方对其的任何修改都不会写回缓存，也不会被后续请求改写。
func (c *Cache) Get(key string) []byte {
	c.mu.RLock()
	defer c.mu.RUnlock()
	snap, ok := c.values[key]
	if !ok {
		return nil
	}
	out := make([]byte, len(snap))
	copy(out, snap)
	return out
}
