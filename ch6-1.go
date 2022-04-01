package main

import "fmt"

func failedUpdate(g *int) {
	x := 10
	g = &x
	fmt.Println(g)
}

func update(px *int) {
	*px = 20
}

func main() {
	var f *int
	failedUpdate(f)
	fmt.Println(f)

	x := 10
	failedUpdate(&x)
	fmt.Println(x)
	update(&x)
	fmt.Println(x)
}
