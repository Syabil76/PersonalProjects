package main

import (
	"fmt"
	"math/rand/v2"
)

var list []int

func initialise(parameter int, rang int) {
	for i := 0; i < parameter; i++ {
		list = append(list, rand.IntN(rang))
	}
	fmt.Println("unsorted: ", list)
}

func sort(efe int, list []int) {
	var temp int
	for i := 0; i < efe-1; i++ {
		for j := 0; j < efe-1; j++ {
			if list[j] > list[j+1] {
				temp = list[j]
				list[j] = list[j+1]
				list[j+1] = temp
			}
		}
	}
}

func main() {
	var para, rang int
	fmt.Print("choose paramter ")
	fmt.Println("choose range ")
	fmt.Scan(&para, &rang)
	initialise(para, rang)
	sort(para, list)
	fmt.Println("sorted: ", list)
}
