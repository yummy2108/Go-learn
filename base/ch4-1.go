package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	samples := []string{"hello", "world!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
}
