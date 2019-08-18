package queue

// go:generate mockgen -destination ./mock/mock_queue.go -package mockqueue queue/queue Queue
// Queue define interface of queue
type Queue interface {
	EnQueue(interface{})
	DeQueue() interface{}
	Head() interface{}
	IsEmpty() bool
	Size() int
}
