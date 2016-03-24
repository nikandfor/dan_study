package main

import "fmt"

func main() {
	var arr [6]int64 // array
	for i := range arr {
		arr[i] = int64(i)
	}

	fmt.Printf("arr : %v\n", arr)

	var s []int64 // slice

	//	s = arr    // NO!!
	s = arr[:]   // OK
	s = arr[2:6] // OK
	// --

	s = arr[2:5:6] // OK // s -> arr from 2 to 6 element, but len = (5-2)

	fmt.Printf("s 1 : %v  - arr[2:5:6] - ссылка на элементы arr с 2 по 6 (до 6), но выводятся только до 5\n", s)

	arr[3] = 50

	fmt.Printf("s 2 : %v\n", s)

	s = append(s, 6) // add element

	arr[4] = 40

	fmt.Printf("s 3 : %v  - добавился 6 и s[2] изменился, потому что s ссылка на arr\n", s)

	s = append(s, 10, 20, 30, 40) // add many elements

	arr[4] = 45

	fmt.Printf("s 4 : %v - s[2] != arr[4] потому что после добавления пришлось скопировать в массив побольше\n", s)

	// можно выделять новые слайсы так
	s2 := make([]int, 10)     // 10 элементов
	s3 := make([]int, 10, 30) // 30 элементов и длина 10 первых из них
	_, _ = s2, s3             // заюзать переменные, чтоб компилятор не ругался
}
