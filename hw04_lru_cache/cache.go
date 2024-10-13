package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if elem, found := c.items[key]; found {
		elem.Value.(*cacheItem).value = value
		c.queue.MoveToFront(elem)
		return true
	}
	if c.queue.Len() >= c.capacity {
		old := c.queue.Back()
		if old != nil {
			oldKey := old.Value.(*cacheItem).key
			c.queue.Remove(old)
			delete(c.items, oldKey)
		}
	}
	freshItem := c.queue.PushFront(&cacheItem{key: key, value: value})
	c.items[key] = freshItem
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if elem, found := c.items[key]; found {
		c.queue.MoveToFront(elem)
		return elem.Value.(*cacheItem).value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}
