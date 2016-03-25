package main

import "fmt"

func main() {
	a := 3

	if a < 5 {
		fmt.Println("a < 5")
	} else {
		fmt.Println("else a < 5")
	}

	switch a {
	case 1, 2:
		fmt.Println("a is: 1, 2")
	case 3, 4:
		fmt.Println("a is: 3, 4")
	case 5:
		fmt.Println("a is: 5")
	default:
		fmt.Println("a is: not 1, 2, 3, 4, 5")
	}

	//	print()
	//	println()
	//	copy()
	//	close()
	//	make()
	//  len()
}
