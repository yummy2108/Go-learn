package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	a := x[1:]
	b := x[1:3]
	c := x[:]
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
