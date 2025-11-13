package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

type Linkedlist struct {
	Head *Node
}

func (l *Linkedlist) arraycon(arr []int) {
	for _, val := range arr {
		newNode := &Node{Data: val}

		if l.Head == nil {
			l.Head = newNode
			continue
		}
		curr := l.Head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = newNode
	}
}

func (l *Linkedlist) Disaply() {
	curr := l.Head
	for curr != nil {
		fmt.Print(curr.Data,"->")
		curr=curr.Next
	}
	fmt.Println("nill")
}
func main() {
list:=&Linkedlist{}
arr:=[]int{10,20,30,40}
list.arraycon(arr)
list.Disaply()
}