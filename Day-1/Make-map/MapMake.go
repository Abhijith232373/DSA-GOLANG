package main

import "fmt"

func mapp() {
	x := new(int)
	*x = 100
	fmt.Println("learn map =",*x)
}


func makee(){
	m:=make([]int,5)
	m[0]=10
	m[1]=20
	fmt.Println("leran make =",m)
}
func main() {
mapp()
makee()
}