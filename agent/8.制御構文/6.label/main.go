package main

import "fmt"

// ラベル付きfor

func main() {
	// Loop:
	//
	//	for {
	//		for {
	//			for {
	//				fmt.Println("START")
	//				break Loop
	//			}
	//			fmt.Println("処理をしない")
	//		}
	//		fmt.Println("処理をしない")
	//	}
	//	fmt.Println("END")

	// continueと組み合わせるパターン
Loop:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if i > 1 {
				continue Loop
			}
			fmt.Println("処理をない")
		}
	}

}
