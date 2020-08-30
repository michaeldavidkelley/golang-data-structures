package queue

type StringQueue interface {
	Enqueue(string)
	Dequeue() string
	Size() int
	IsEmpty() bool
}

func NewStringQueue() StringQueue {
	return &stringQueue{}
}

type stringQueue struct {
	StringQueue

	queue []string
}

func (q *stringQueue) Enqueue(data string) {
	q.queue = append(q.queue, data)
}

func (q *stringQueue) Dequeue() string {
	size := q.Size()

	if size == 0 {
		return ""
	}

	data := q.queue[size-1]
	q.queue = q.queue[:size-1]

	return data
}

func (q *stringQueue) Size() int {
	return len(q.queue)
}

func (q *stringQueue) IsEmpty() bool {
	return q.Size() == 0
}
