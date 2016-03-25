package main

import (
	"flag"
	"fmt"
)

var (
	N      = flag.Int("N", 100, "number of numbers")
	M      = flag.Int("M", 4, "digit to search")
	silent = flag.Bool("silent", false, "dont print hint messages")
)

func main() {
	flag.Parse()

	ncount := [20]int{}
	count := 0
	for i := 0; i < *N; i++ {
		q := 0
		add := false

		for j := 1; j <= i; j *= 10 {
			a := (i / j) % 10
			if a == *M {
				add = true
				ncount[q]++
			}

			q++
		}

		if i == 0 && *M == 0 {
			add = true
			ncount[q]++
		}

		if add {
			count++
			if !*silent {
				fmt.Printf("i %v   count++\n", i)
			}
		}
	}

	fmt.Printf("count: %v  %v\n", count, ncount)

	_ = count
}
