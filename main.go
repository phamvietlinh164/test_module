package main

import (
	"fmt"
	"sort"
)

func numCar(a []int, b []int) int {
	var people int
	var car int
	for _, v := range a {
		people = people + v
	}
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})
	for _, v := range b {
		people = people - v
		car = car + 1
		if people <= 0 {
			break
		}
	}
	return car
}

func main() {
	a := numCar([]int{4, 1, 2, 2, 2}, []int{4, 5, 5, 6, 5})
	fmt.Println(a)
}
