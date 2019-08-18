package queue

import (
	"sync"
)

type queue struct {
	items *[]interface{}
	lock  sync.RWMutex
}

// New creates a new Queue
func New() Queue {
	items := new([]interface{})
	return &queue{
		items: items,
	}
}

// EnQueue adds an Item to the end of the queue
func (q *queue) EnQueue(item interface{}) {
	q.lock.Lock()
	*q.items = append(*q.items, item)
	q.lock.Unlock()
}

// DeQueue removes an Item from the start of the queue
func (q *queue) DeQueue() interface{} {
	q.lock.Lock()
	item := (*q.items)[0]
	*q.items = (*q.items)[1:len(*q.items)]
	q.lock.Unlock()
	return item
}

// Head returns the item next in the queue, without removing it
func (q *queue) Head() interface{} {
	if q.Size() == 0 {
		return nil
	}

	q.lock.RLock()
	item := (*q.items)[0]
	q.lock.RUnlock()
	return item
}

// IsEmpty returns true if the queue is empty
func (q *queue) IsEmpty() bool {
	return len(*q.items) == 0
}

// Size returns the number of Items in the queue
func (q *queue) Size() int {
	return len(*q.items)
}
