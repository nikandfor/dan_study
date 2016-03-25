package main

import "fmt"

func main() {
	ncount := [20]int{}
	count := 0
	for i := 0; i < 100; i++ {
		a := (i / 1) % 10
		b := (i / 10) % 10

		if a == 5 || b == 5 {
			count++
			fmt.Printf("i %v   count++\n", i)
		}

		if a == 5 {
			ncount[0]++
		}
		if b == 5 {
			ncount[1]++
		}
	}

	fmt.Printf("count: %v  %v\n", count, ncount)
}
