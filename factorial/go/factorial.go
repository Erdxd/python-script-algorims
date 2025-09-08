package main

import "fmt"

func main() {
	fmt.Println(factorial(11))
}
func factorial(a int) int {
	z := 0
	p := 1

	for z < a {
		z++
		p *= z

	}
	return p

}
