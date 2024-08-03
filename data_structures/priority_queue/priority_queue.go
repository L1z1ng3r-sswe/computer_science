package priority_queue

type PriorityQueueImpl interface {
	Less(i, j int) bool
	Equal(i, j int) bool
	Empty() bool

	Length() int
	Swap(i, j int)
	Push(data interface{}) interface{}
	Pop() interface{}
}

type PriorityQueue interface {
	Enqueue(interface{})
	Dequeue() interface{}
	down(i int)
	up(i int)
}

type priorityQueue struct {
	impl PriorityQueueImpl
}

func New(impl PriorityQueueImpl) *priorityQueue {
	return &priorityQueue{impl: impl}
}

func (pq *priorityQueue) Enqueue(value interface{}) {
	pq.impl.Push(value)
	pq.up(pq.impl.Length() - 1)
}

func (pq *priorityQueue) Dequeue() interface{} {
	if pq.impl.Empty() {
		return nil
	}

	pq.impl.Swap(0, pq.impl.Length()-1)
	minValue := pq.impl.Pop()
	pq.down(0)

	return minValue
}

func (pq *priorityQueue) up(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if pq.impl.Less(index, parentIndex) {
			pq.impl.Swap(index, parentIndex)
			index = parentIndex
		} else {
			break
		}
	}
}

func (pq *priorityQueue) down(index int) {
	length := pq.impl.Length()
	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		smallest := index

		if leftChild < length && pq.impl.Less(leftChild, smallest) {
			smallest = leftChild
		}
		if rightChild < length && pq.impl.Less(rightChild, smallest) {
			smallest = rightChild
		}
		if smallest == index {
			break
		}

		pq.impl.Swap(index, smallest)
		index = smallest
	}
}
