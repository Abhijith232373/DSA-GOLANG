package main

import "fmt"

func rev(arr []int) {
	for i, j := 0, len(arr)-1;i < j;i,j=i+1,j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	num := []int{1, 2, 43, 4, 5}
	rev(num)
	fmt.Println(num)
}