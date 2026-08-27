package leasecatalog

type Cache struct { values map[string][]byte }

func NewCache() *Cache { return &Cache{values: make(map[string][]byte)} }

// Put 存入时拷贝 value，缓存持有自己的独立副本，
// 调用方后续改写源切片不会泄漏进缓存。
func (c *Cache) Put(key string, value []byte) {
	cp := make([]byte, len(value))
	copy(cp, value)
	c.values[key] = cp
}

// Get 返回独立副本，改写返回值不能反向污染缓存。
func (c *Cache) Get(key string) []byte {
	v, ok := c.values[key]
	if !ok {
		return nil
	}
	out := make([]byte, len(v))
	copy(out, v)
	return out
}
