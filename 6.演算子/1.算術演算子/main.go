package main

import "fmt"

// 算術演算子

func main() {
	fmt.Println(1 + 1)

	fmt.Println("ABC" + "DEF")
	fmt.Println(5 - 1)
	fmt.Println(5 * 2)
	fmt.Println(10 / 2)

	fmt.Println(10 % 3)

	n := 0
	n += 4
	fmt.Println(n)
	n ++
	fmt.Println(n)
	n --
}