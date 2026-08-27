package segmentcatalog

// Cache stores decoded segment tag bytes keyed by tag.
type Cache struct {
	values map[string][]byte
}

// NewCache returns an empty Cache.
func NewCache() *Cache { return &Cache{values: make(map[string][]byte)} }

// Put stores an isolated copy of value, so later mutations of value by the
// caller cannot rewrite the cached bytes. This is what stops a cached first
// batch from silently reflecting a second batch decoded into the same
// backing array.
func (c *Cache) Put(key string, value []byte) {
	stored := make([]byte, len(value))
	copy(stored, value)
	c.values[key] = stored
}

// Get returns an isolated copy of the cached bytes, so reader-side writes
// cannot pollute the cache (and cannot leak into other batches through a
// shared backing array). Returns nil for unknown keys.
func (c *Cache) Get(key string) []byte {
	stored, ok := c.values[key]
	if !ok {
		return nil
	}
	out := make([]byte, len(stored))
	copy(out, stored)
	return out
}
