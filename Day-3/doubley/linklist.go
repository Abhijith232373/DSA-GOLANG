package main

import "fmt"

type Node struct{
	Data int
	Prev *Node
	Next *Node
}

type DoublyLinkedList struct{
	Head *Node
}

func (l *DoublyLinkedList)inserFront(val int){
	newNode:=&Node{Data:val}

	if l.Head == nil{
		l.Head=newNode
		return
	}
	newNode.Next=l.Head
	l.Head.Prev=newNode
	l.Head=newNode
}

func(l *DoublyLinkedList)Deletelast(){
	if l.Head==nil{
		return
	}
	if l.Head.Next==nil{
		l.Head=nil
		return
	}

	curr:=l.Head
	for curr.Next!=nil{
		curr=curr.Next
	}
	curr.Prev.Next=nil
}


func (l *DoublyLinkedList)Disaply(){
	curr:=l.Head
	for curr!=nil{
		fmt.Print(curr.Data,"<->")
		curr=curr.Next
	}
	fmt.Println("nil")
}
func main(){
	list:=&DoublyLinkedList{}
	list.inserFront(10)
	list.inserFront(20)
	list.inserFront(30)

	list.Disaply()

	list.Deletelast()
	list.Disaply()

}