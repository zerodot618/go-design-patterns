package strategy

type cache struct {
	storage       map[string]string
	evicationAlgo evictionAlgo
	capacity      int
	maxCapacity   int
}

func initCache(e evictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:       storage,
		evicationAlgo: e,
		capacity:      0,
		maxCapacity:   2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evicationAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evicationAlgo.evict(c)
	c.capacity--
}
