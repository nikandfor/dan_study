package main

import "fmt"

func main() {
	//	var a int = 2
	//	b := 4

	fmt.Printf("inf loop\n")
	i := 0
	for {
		fmt.Printf("i1 : %v\n", i)
		i++

		if !(i < 3) {
			break
		}
	}

	fmt.Printf("2 loop\n")
	i = 0
	for i < 3 {
		fmt.Printf("i2 : %v\n", i)
		i++
	}

	fmt.Printf("3 loop\n")
	for i = 0; i < 3; i++ {
		if i == 1 {
			continue
		}
		fmt.Printf("i3 : %v\n", i)
	}
}
