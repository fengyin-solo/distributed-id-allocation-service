package zonecatalog

type Cache struct{ values map[string][]byte }

func NewCache() *Cache { return &Cache{values: make(map[string][]byte)} }

// Put 写入区域值。缓存内部持有的是 value 的副本，调用方之后再修改 value 不会影响缓存。
func (c *Cache) Put(key string, value []byte) {
	stored := make([]byte, len(value))
	copy(stored, value)
	c.values[key] = stored
}

// Get 读取区域值。返回缓存内部数据的副本，调用方修改返回值不会影响缓存。
func (c *Cache) Get(key string) []byte {
	v, ok := c.values[key]
	if !ok {
		return nil
	}
	out := make([]byte, len(v))
	copy(out, v)
	return out
}
