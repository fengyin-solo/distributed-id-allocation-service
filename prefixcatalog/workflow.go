package prefixcatalog

type Cache struct { values map[string][]byte }

func NewCache() *Cache { return &Cache{values: make(map[string][]byte)} }

// Put stores an independent copy of value so later mutation of the caller's
// slice (or a reused buffer backing it) cannot corrupt the cached entry.
func (c *Cache) Put(key string, value []byte) {
	stored := make([]byte, len(value))
	copy(stored, value)
	c.values[key] = stored
}

// Get returns an independent copy of the cached entry so callers cannot mutate
// the cached content through the returned slice.
func (c *Cache) Get(key string) []byte {
	v, ok := c.values[key]
	if !ok {
		return nil
	}
	out := make([]byte, len(v))
	copy(out, v)
	return out
}
