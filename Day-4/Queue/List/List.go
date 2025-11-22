package main

import "fmt"

type Node struct {
	data int
	next *Node
}
type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) enqueue(val int) {
	newNode := &Node{data: val}

	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
		return
	}
	q.rear.next = newNode
	q.rear = newNode
}

func (q *Queue) deQueue() int {
	if q.front == nil {
		fmt.Println("queue is empty")
		return -1
	}
	val:=q.front.data
	q.front=q.front.next
	

	if  q.front==nil{
		q.rear=nil
	}
	return val
}
func (q *Queue)peek()int{
	if q.front==nil{
		fmt.Println("queue empty")
		return -1
	}
	return q.front.data
}

func (q *Queue)Disaply(){
	if q.front==nil{
		fmt.Println("queue is  empty")
		return
	}
	temp:=q.front
	for temp!=nil{
		fmt.Print(temp.data," ")
		temp=temp.next
	}
	fmt.Println()
}
func main() {
q:=&Queue{}
q.enqueue(50)
q.enqueue(40)
q.enqueue(30)
q.enqueue(1)
q.Disaply()
q.deQueue()
q.Disaply()
fmt.Println(q.peek())
}