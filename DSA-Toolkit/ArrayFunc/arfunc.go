package main

func innsert( arr[]int,index int, val int)[]int{
	arr=append(arr, 0)
	for i:=len(arr)-1;i>index;i--{
		arr[i]=arr[i-1]
	}
	arr[index]=val
	return arr
}
func Delete (arr []int,index int)[]int{
	return append(arr[:index],arr[index+1:]...)
}

func search(arr []int,val int)int{
    for i, v := range arr {
        if v == val {
            return i
        }
    }
    return -1
}