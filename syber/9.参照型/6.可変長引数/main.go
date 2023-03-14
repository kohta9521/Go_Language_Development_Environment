package main

import "fmt"

func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	fmt.Println(Sum(1, 2, 3))

	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	sl := []int{1, 2, 3}
	fmt.Println(Sum(sl...))
}
