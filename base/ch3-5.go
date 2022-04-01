package main

import "fmt"

func main() {
	intSet := map[int]bool{}
	vals := []int{1, 2, 3, 4, 5, 5, 6, 7}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
}
