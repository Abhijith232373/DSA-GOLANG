package main

import "fmt"

type Stack struct {
	arr [6]int
	top int
}

func (s *Stack) push(val int) {
	if s.top==len(s.arr)-1{
		fmt.Println("stack is full")
		return
	}
	s.top++
	s.arr[s.top]=val
}
func (s *Stack) pop() int {
	if s.top == -1 {
		fmt.Println("stack is empty")
		return -1
	}
	val:=s.arr[s.top]
	s.top--
	return val
}
func (s *Stack)peek()int{
	if s.top==-1{
		fmt.Println("stack emty")
		return -1
	}
	return s.arr[s.top]
}
func (s *Stack)Display(){
	if s.top==-1{
		fmt.Println("stack is empty")
		return 
	}
	for i:=s.top;i>=0;i--{
		fmt.Print(s.arr[i]," ")

	}
	fmt.Println()
}

func main() {
s:=&Stack{top: -1}
s.push(50)
s.push(40)
s.push(30)
s.push(1)
s.Display()
fmt.Println(s.pop())
s.Display()
fmt.Println(s.peek())
s.Display()
}