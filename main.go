package main

import (
	"fmt"
)

func main() {
	fmt.Println(Sum(1 ,2))
}

func Sum(a int, b int) int {
	sum := a + b
	return sum
}