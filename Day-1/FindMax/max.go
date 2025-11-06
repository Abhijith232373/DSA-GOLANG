package main

import "fmt"

func main() {
	arr := new([5]int)
	arr[0] = 10
	arr[1] = 1
	arr[2] = 110
	arr[3] = 13
	arr[4] = 20
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	fmt.Println("Max value is :",max)

}