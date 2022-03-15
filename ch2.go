package main

import "fmt"

func main() {
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x)
	fmt.Println(y)

	var a int = 10
	var b float64 = 30.2
	var c float64 = float64(a) + b
	var d int = a + int(b)
}
