package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
	Len() int
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c lruCache) Set(key Key, value interface{}) bool {
	item, ok := c.items[key]
	if !ok {
		if c.capacity > c.queue.Len() {
			n := c.queue.PushFront(value.(int))
			c.items[key] = n
			return false
		}
		del := c.queue.Back()
		for k, v := range c.items {
			if v == del {
				delete(c.items, k)
			}
		}
		c.queue.Remove(del)
		n := c.queue.PushFront(value.(int))
		c.items[key] = n
		return false
	}
	item.Value = value.(int)
	c.queue.MoveToFront(item)
	return true
}

func (c lruCache) Get(key Key) (interface{}, bool) {
	item, ok := c.items[key]
	if !ok {
		return nil, false
	}
	c.queue.MoveToFront(item)
	return c.items[key].Value.(int), true
}

func (c lruCache) Clear() {
	c.queue.RemoveAll()
}

func (c lruCache) Len() int {
	return c.queue.Len()
}
