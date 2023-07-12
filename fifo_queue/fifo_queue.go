package fifo_queue

import "fmt"

type Queue interface {
	Add(link string)
	Get() string
	IsEmpty() bool
}

type queueImpl struct {
	tasks chan string
}

func New() Queue {
	return &queueImpl{
		tasks: make(chan string, 100000),
	}
}

func (q *queueImpl) Add(link string) {
	fmt.Println("adding  to queue ", len(q.tasks))
	select {
	case q.tasks <- link:
	}
	fmt.Println("added  to queue ", link)
}

func (q *queueImpl) Get() string {
	if q.IsEmpty() {
		return ""
	}
	return <-q.tasks
}

func (q *queueImpl) IsEmpty() bool {
	return len(q.tasks) == 0
}
