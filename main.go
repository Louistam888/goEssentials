package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Printf("hello  name")
	// var agency string = "Fast Track"
	// fmt.Println("welcome to", agency)
	// name := "Talal" //type inference
	// fmt.Printf("hellot, %v\n", name, )

	// var totalCars int = 50
	// fmt.Println(totalCars)
	// startingPrice := 29.99
	// fmt.Printf("our prices start at %v\n", startingPrice)
	// fmt.Println("Take yourpick")

	// insuranceIncluded := true
	// fmt.Println("Prices included", insuranceIncluded)

	str := "hello, world"
	length := len(str)
	fmt.Println(length)

	str1 := "Go programming"
	str2 := "go programming"

	fmt.Println(strings.EqualFold(str1, str2))

	str4 := "hello, World"
	wIndex := strings.Index(str4,"W") //index 7
	fmt.Println(wIndex)

	substr := str4[wIndex:12] // starting atindex 7
	fmt.Println(substr)
}
