package main

import "fmt"

type Queue struct {
	arr   [5]int
	front int
	rear  int
}

func (q *Queue) enqueue(val int) {
	if q.rear == len(q.arr) {
		fmt.Println("queue is bull")
		return
	}
	q.arr[q.rear] = val
	q.rear++
}
func (q *Queue) deQueue() int {
	if q.front == q.rear {
		fmt.Println("queueue is empty")
		return -1
	}
	val := q.arr[q.front]
	q.front++
	return val

}

func (q *Queue) Disaply() {
	if q.front == q.rear {
		fmt.Println("queue is empty")
		return
	}
	for i := q.front; i < q.rear; i++ {
		fmt.Print(q.arr[i], " ")

	}
	fmt.Println()

}

func main() {
	q := &Queue{}
	q.enqueue(15)
	q.enqueue(13)
	q.enqueue(14)
	q.enqueue(12)
	q.enqueue(11)
	q.Disaply()
	fmt.Println(q.deQueue())
	q.Disaply()
}
