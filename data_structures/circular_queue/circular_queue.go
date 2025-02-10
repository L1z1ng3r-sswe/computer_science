package circular_queue

import (
	"errors"
	"fmt"
	"sync"
)

type CircularQueue struct {
	data     []interface{}
	cap      int
	currSize int
	front    int // the element to return first
	rear     int // always empty
	mu       sync.RWMutex
	notEmpty *sync.Cond
	notFull  *sync.Cond
}

func New(buffSize int) *CircularQueue {
	cq := &CircularQueue{
		data: make([]interface{}, buffSize),
		cap:  buffSize,
	}

	cq.notEmpty = sync.NewCond(&cq.mu)
	cq.notFull = sync.NewCond(&cq.mu)
	return cq
}

func (cq *CircularQueue) Enqueue(newData interface{}) interface{} {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	for cq.isFull() {
		fmt.Println("Waiting for space to free up...")
		cq.notFull.Wait()
	}

	cq.data[cq.rear] = newData
	cq.rear = (cq.rear + 1) % cq.cap
	cq.currSize++
	cq.notEmpty.Broadcast()

	return newData
}

func (cq *CircularQueue) Dequeue() interface{} {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	for cq.isEmpty() {
		fmt.Println("Waiting for buffer to fill in...")
		cq.notEmpty.Wait()
	}

	res := cq.data[cq.front]
	cq.data[cq.front] = nil // if u want to save up space and use visualization.
	cq.front = (cq.front + 1) % cq.cap
	cq.currSize--
	cq.notFull.Broadcast()

	return res
}

func (cq *CircularQueue) ObserveFront() (interface{}, error) {
	cq.mu.RLock()
	defer cq.mu.RUnlock()

	if cq.isEmpty() {
		return nil, errors.New("Buffer Is Empty")
	}

	return cq.data[cq.front], nil
}

func (cq *CircularQueue) ObserveLast() (interface{}, error) {
	cq.mu.RLock()
	defer cq.mu.RUnlock()

	if cq.isEmpty() {
		return nil, errors.New("Buffer Is Empty")
	}

	lastIndex := (cq.rear - 1 + cq.cap) % cq.cap
	return cq.data[lastIndex], nil
}

func (cq *CircularQueue) Visualization() string { // [1, _, _, _, _, _, _, _, _, _]
	var visualization = "["
	for i := 0; i < cq.cap; i++ {
		if cq.data[i] != nil {
			serialized := fmt.Sprintf("%v", cq.data[i])
			visualization += serialized
		} else {
			visualization += "_"
		}
		if i+1 < cq.cap {
			visualization += ","
		}
	}

	return visualization + "]"
}

func (cq *CircularQueue) isEmpty() bool {
	if cq.currSize <= 0 { // <= is more resilient than ==
		return true
	}

	return false
}

func (cq *CircularQueue) isFull() bool {
	if cq.currSize >= cq.cap { // >= is more resilient than ==
		return true
	}

	return false
}
