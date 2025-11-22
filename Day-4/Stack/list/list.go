package main

import "fmt"

type Node struct {
	Next *Node
	data int
}
type Linkedlist struct {
	top *Node
}

func (l *Linkedlist) push(val int) {
	newNode := &Node{data: val}
	newNode.Next = l.top
	l.top = newNode

}

func (l *Linkedlist) pop() int {
	if l.top == nil {
		fmt.Println("stack is empty")
		return -1
	}
	val:=l.top.data
	l.top=l.top.Next
	return val
}
func (l *Linkedlist)peek()int{
	if l.top==nil{
		fmt.Print("stack is empty")
		return -1
	}
	return l.top.data
}


func (l *Linkedlist)Disaply(){
	if l.top==nil{
		fmt.Println("stack is empty")
		return
	}
	// fmt.Println("stack")
	temp:=l.top
	for temp!=nil{
		fmt.Print(temp.data," ")
		temp=temp.Next
	}
	fmt.Println()
}
func main() {
s:=&Linkedlist{}
s.push(10)
s.push(30)
s.push(20)
s.push(1)
s.Disaply()
fmt.Println(s.pop())
s.Disaply()
fmt.Println(s.peek())

}