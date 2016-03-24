package main

import "fmt"

func main() {
	s2 := make([]int, 10)
	for i := range s2 {
		s2[i] = i
	}

	fmt.Printf("s2 1 : %v\n", s2)

	fmt.Printf("копируем %v в %v...  влево пересекающиеся куски\n", s2[1:4], s2[2:5])
	copy(s2[2:5], s2[1:4])

	fmt.Printf("s2 2 : %v\n", s2)

	// set values 0..9
	for i := range s2 {
		s2[i] = i
	}

	fmt.Printf("копируем %v в %v...  вправо пересекающиеся куски\n", s2[2:], s2[1:4])
	copy(s2[1:4], s2[2:])

	fmt.Printf("s2 3 : %v\n", s2)
}
