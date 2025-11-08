package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("orginal slice:",sl)
	// ........................................insert
	index := 2
	value := 10
	sl = append(sl[:index], append([]int{value}, sl[index:]...)...)
	fmt.Println("added 10 in 2nd index position:",sl)
	// .......................................................................delete
	del:=1
	sl=append(sl[:del],sl[del+1:]...)
	fmt.Println("After 1nd index delete:",sl)
}