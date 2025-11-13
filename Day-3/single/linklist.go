package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type Linkedlist struct {
	Head *Node
}

func (l *Linkedlist)insert(val int) {
	newNode := &Node{Data: val}

	if l.Head == nil {
		l.Head = newNode
		return
	}
	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
}

func (l *Linkedlist) Delete(val int) {
	if l.Head == nil {
		fmt.Println("Empty list")
		return
	}
	if l.Head.Data==val{
		l.Head=l.Head.Next
		return
	}
	curr:=l.Head
	    for curr.Next!=nil&& curr.Next.Data!=val{
		curr=curr.Next
	}
	if curr.Next==nil{
		fmt.Println("value not found")
		return
	}
	curr.Next=curr.Next.Next
}

func (l *Linkedlist)search (val int)bool{
	curr:=l.Head
	for curr !=nil{
		if curr.Data==val{
			return true
		}
		curr=curr.Next
	}
	return  false
}

func (l *Linkedlist)Disaply(){
	curr:=l.Head
	for curr !=nil{
		fmt.Print(curr.Data,"->")
		curr=curr.Next
	}
	fmt.Println("nil")
}

func main(){
	list := &Linkedlist{}

	list.insert(10)
	list.insert(14)
	list.insert(15)
	list.insert(13)
	list.insert(12)
	list.Disaply()

	fmt.Println("Search 20:",list.search(10))
	fmt.Println("Search 20:",list.search(15))
	list.Delete(12)
	list.Disaply()
}