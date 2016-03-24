package main

import "fmt"

func main() {
	var arr [6]int64

	arr[4] = 40

	b := arr

	arr[2] = 20

	var c *[6]int64 = &arr // pointer to array

	arr[0] = 10

	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("b  : %v  - b[4] != 40, потому что присвоили arr[4] - b копия\n", b)
	fmt.Printf("c  : %v  - c[0] = 10, но мы присвоили arr[0] - c указатель на arr\n", c)
}
