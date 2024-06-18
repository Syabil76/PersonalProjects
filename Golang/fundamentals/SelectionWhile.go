package main

import "fmt"

func main() {
	var choice string
	fmt.Printf("Choose A or B (Case sensitive)")
	for {
		fmt.Scanf("%s", &choice)
		if choice == "A" {
			fmt.Printf("The choice is %s", choice)
			break
		} else if choice == "B" {
			fmt.Printf("The choice is %s", choice)
			break
		}
	}
}
