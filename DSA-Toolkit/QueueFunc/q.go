package main

import "fmt"

type Queue struct {
	arr   [5]int
	front int
	rear  int
}

func (q *Queue) enqueue(val int) {
	if q.rear == len(q.arr) {
		fmt.Println("queue is full")
		return
	}
	q.arr[q.rear]=val
	q.rear++
}


func (q *Queue)dequeue()int{
	if q.front==q.rear{
		fmt.Println("qusue is empty")
		return  -1
	}
	val:=q.arr[q.front]
	q.front++
	return val
}
