package main

import (
	"Assignment/camelCase"
	"Assignment/fileHandling"
	"Assignment/htmlParser"
	"fmt"
)

func main() {
	fmt.Println("******* Exercise 1 *******")
	fmt.Println("Number of words:", camelCase.NumberOfWords("this is nice"))
	fmt.Println("__________________________")

	fmt.Println("******* Exercise 2 *******")
	fileHandling.Run()
	fmt.Println("__________________________")

	fmt.Println("******* Exercise 3 *******")
	fmt.Println(htmlParser.ParseHtml("ex4.html"))
	fmt.Println("__________________________")
}