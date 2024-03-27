package main

import "fmt"

func main() {
	i := 0
	for i <= 100 {
		if i%5 == 0 && i%7 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("fizz")
		} else if i%7 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
		i = i + 1
	}
}
