package main

import (
	"fmt"
)

type Stack struct {
	arr []int
}

func (s *Stack) push(val int) {
	s.arr=append(s.arr, val)
	}

func (s *Stack)pop()int{
	if s.arr==nil{
		fmt.Println("stack is empty")
		return -1
	}
	top:=s.arr[len(s.arr)-1]
	s.arr=s.arr[:len(s.arr)-1]
	return  top
}

func (s *Stack)peek()int{
	if s.arr==nil{
		fmt.Println("stack is empty")
		return -1
	}
	return s.arr[len(s.arr)-1]
}