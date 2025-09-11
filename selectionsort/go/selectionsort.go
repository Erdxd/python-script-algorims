package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 100, 5}
	fmt.Println(selectionsort(arr))
}
func findsmallest(arr []int) int {
	smallest := arr[0]
	smallest_index := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}
	return smallest_index
}
func selectionsort(arr []int) []int {
	NewArr := []int{}
	copiedArr := make([]int, len(arr))
	copy(copiedArr, arr)
	for i := 0; i < len(arr); i++ {
		smallest := findsmallest(copiedArr)
		NewArr = append(NewArr, copiedArr[smallest])
		copiedArr = append(copiedArr[:smallest], copiedArr[smallest+1:]...)

	}
	return NewArr
}
