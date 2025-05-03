package main

import (
	"fmt"
)

func main() {
	var vvod string
	var a int
	var b int
	var s int
	fmt.Println("vvedite deistvie")
	x, y, z := fmt.Scan(&vvod, &a, &b)
	if x == "add" {
		s = add(y, z)
	}
	if x == "sub" {
		s = sub(y, z)
	}
	if x == "mul" {
		s = mul(y, z)
	}
	if x == "div" {
		s = div(y, z)
	}
	fmt.Println(s)
}
