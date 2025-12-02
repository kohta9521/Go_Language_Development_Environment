package main

import "fmt"

// 変数

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	fmt.Println("== 明示的な定義 ================")
	// 明示的な定義
	var i int = 100
	fmt.Println(i)

	// 文字列型の変数定義
	var s string = "Hello Golang" // 文字列型はダブルクオーテーションで囲むことで表現可能
	fmt.Println(s)

	// 複数の変数で型が同じ場合はまとめて型指定が可能
	var t, f bool = true, false
	fmt.Println(t, f)

	// 異なる型の変数をまとめて定義する
	var (
		i2 int = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	// 変数に初期値を設定しない場合
	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	fmt.Println("== 暗黙的な定義 ================")

	// 暗黙的な定義
	i4 := 400
	fmt.Println(i4)

	// i4 = "Hello, "
	// fmt.Println(i4)

	outer()

}