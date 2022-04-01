package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	people := []Person{
		{"A", "C", 1},
		{"B", "B", 2},
		{"C", "A", 3},
	}
	fmt.Println(people)

	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})

	fmt.Println(people)
}
