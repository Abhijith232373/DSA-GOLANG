package main

import (
	"fmt"
)

type Node struct {
	Next *Node
	data int
}
type Linkedlist struct {
	Head *Node
}

func (l *Linkedlist) insertHead(val int) {
	newNode := &Node{data: val}
	if l.Head != nil {
		newNode.Next = l.Head
		l.Head = newNode
	}
}

func (l *Linkedlist) insertTail(val int) {
	newNode := &Node{data: val}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	temp := l.Head
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = newNode
}

func (l *Linkedlist) Delete(val int){
	if l.Head == nil {
		fmt.Println("stack empty")
		return 
	}
	if l.Head.data==val{
		l.Head=l.Head.Next
		return
	}
	temp:=l.Head
	for temp.Next!=nil&&temp.Next.data!=val{
		temp=temp.Next

	}

	if temp.Next==nil{
		fmt.Println("value not founded")
		return
	}
	temp.Next=temp.Next.Next


}
