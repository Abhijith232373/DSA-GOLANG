package main

import "fmt"

func merge(num1 []int, m int, num2 []int, n int){
	i := m - 1
	j := n - 1
	k := m + n - 1
	for j >= 0 {
		if i >= 0 && num1[i] > num2[j] {
			num1[k] = num1[i]
			i--
		} else {
			num1[k] = num2[j]
			j--
		}
		k--
	}

}
func main() {
	num1:=[]int{1,2,3,0,0,0}
	merge(num1,3,[]int{2,5,6},3)
	fmt.Println(num1)
}