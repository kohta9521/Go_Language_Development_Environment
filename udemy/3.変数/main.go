package main

import "fmt"

// 変数

func main() {
	// 明示的な定義
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go lang!!!!"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "Go !!!"
	)
	fmt.Println(i2, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3)
}
