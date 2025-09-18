package main

import "fmt"

func main() {
	d := []int{1, 2, 3, 4, 5}
	fmt.Println(sumwithrecursia(d))
}
func sumwithrecursia(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + sumwithrecursia(arr[1:])
}
