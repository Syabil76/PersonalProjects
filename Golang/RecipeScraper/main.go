package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("No arguments detected")
	}

	//collect arguments from user in CLI
	url := os.Args[1]
	collector := colly.NewCollector()

	//Print error messages for debugging purposes- got online
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", url)
	})
	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Entering", url)
	})
	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println("Could not load", e)
	})

	collector.OnHTML(".div-main", func(e *colly.HTMLElement) {

	})

}

type Dictionary map[string]string

type RecipeOverview struct {
	Course, Cuisine, Servings, Calories string
}

type Recipe struct {
	url, dish      string
	ingridients    []Dictionary
	specifications RecipeOverview
}
