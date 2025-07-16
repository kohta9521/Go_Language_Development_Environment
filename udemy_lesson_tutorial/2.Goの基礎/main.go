package main

import (
	"fmt"
	"time"
)

// hello world print function
func main() {
	fmt.Println("Hello, World!")
	t := time.Now()
	fmt.Println(t)
	fmt.Println(time.Now())
}