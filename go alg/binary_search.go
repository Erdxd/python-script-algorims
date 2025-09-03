package main

import "fmt"

func main() {
	d := []int{1, 2, 3, 4, 5}

	fmt.Println(bin_search(d, 3))
}
func bin_search(d []int, number int) int {
	low := 0
	max := len(d) - 1
	for low <= max {
		mid := (low + max) / 2
		guess := d[mid]
		if guess == number {
			return mid
		} else if guess > number {
			max = mid - 1
		} else if guess < number {
			low = mid + 1
		}
	}
	return 0
}
