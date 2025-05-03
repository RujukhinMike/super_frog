package main

import (
	"fmt"
)

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	if b == 0 {
		fmt.Println("Деление на ноль!")
		return 0
	}
	return a / b
}

func main() {}
