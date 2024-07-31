package main

import (
	"fmt"
	"time"

	"github.com/L1z1ng3r-sswe/computer_science/data_structures/circular_queue"
)

func main() {
	cq := circular_queue.New(10)

	go overflower(cq)
	go reader(cq)

	select {}
}

func overflower(cq *circular_queue.CircularQueue) {
	var i int
	for {
		cq.Enqueue(i)
		fmt.Printf("Enqueued: %d\n", i)
		fmt.Println(cq.Visualization())
		i++
		time.Sleep(time.Second * 2)
	}
}

func reader(cq *circular_queue.CircularQueue) {
	for {
		val := cq.Dequeue()
		fmt.Printf("Dequeued: %v\n", val)
		fmt.Println(cq.Visualization())
		time.Sleep(time.Second * 2)
	}
}
