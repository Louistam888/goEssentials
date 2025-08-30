package main

import "fmt"

// func main() {
// 	// fmt.Printf("hello  name")
// 	// var agency string = "Fast Track"
// 	// fmt.Println("welcome to", agency)
// 	// name := "Talal" //type inference
// 	// fmt.Printf("hellot, %v\n", name, )

// 	// var totalCars int = 50
// 	// fmt.Println(totalCars)
// 	// startingPrice := 29.99
// 	// fmt.Printf("our prices start at %v\n", startingPrice)
// 	// fmt.Println("Take yourpick")

// 	// insuranceIncluded := true
// 	// fmt.Println("Prices included", insuranceIncluded)

// 	str := "hello, world"
// 	length := len(str)
// 	fmt.Println(length)

// 	str1 := "Go programming"
// 	str2 := "go programming"

// 	fmt.Println(strings.EqualFold(str1, str2))

// 	str4 := "hello, World"
// 	wIndex := strings.Index(str4,"W") //index 7
// 	fmt.Println(wIndex)

// 	substr := str4[wIndex:12] // starting atindex 7
// 	fmt.Println(substr)
// }

// func main() {
// 	str4 := "hello, World"
// 	wIndex := strings.Index(str4, "W")
// 	fmt.Println(str4[:wIndex])
// 	fmt.Println(str4[wIndex:])
// }

// func main() {
// 	str5 := "Hello Go!"
// 	fmt.Println(strings.Replace(str5, "Go", "World", 1)) //1 says replace the first instance of go
// }

//	func main() {
//		str6:= "Go Programming"
//		fmt.Println(strings.ToUpper(str6))
//		fmt.Println(strings.ToLower(str6))
//	}

// func main() {
// 	str6 := "Go Programming"
// 	fmt.Println(strings.Contains(str6, "Go"))
// 	fmt.Println(strings.Contains(str6, "Golang"))
// }

// func main() {
// 	var tinyUInt uint8 = 255 //unsigned, can only be used for posiitive ints, 8 bits of memory max value can be 255
// }

// TEMP CONVERSION C TO K
func main() {
	// temperatureC := 32
	// temperatureK := 0.0

	// temperatureK = temperatureC + 273.15 // this will throw error as we are adding float with int

	// SOLUTION 1 typecast
	// temperatureK = float64(temperatureC) + 273.15

	// SOLUTION 2 change vars beforeadding
	temperatureC2 := 32.0
	temperatureK2 := 0.0

	temperatureK2 = temperatureC2 + 273.15

	fmt.Println("Temperature in C", temperatureC2)
	fmt.Println("Temperature in K", temperatureK2)

}
