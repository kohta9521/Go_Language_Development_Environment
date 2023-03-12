package main

import "fmt"

// Switch文

func main() {
	// n := 5
	// switch n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("i don 't  knoe")
	// }

	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("i don 't knoe")
	}

	// n := 6
	// switch {
	// case n > 0 && n < 4:
	// 	fmt.Println("0 < n < 4")
	// case n > 3 && n < 7:
	// 	fmt.Println("3 < n < 7")
	// }

	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("i don 't knoe")
	}
}
