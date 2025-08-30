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
	// temperatureC2 := 32.0
	// temperatureK2 := 0.0

	// temperatureK2 = temperatureC2 + 273.15

	// fmt.Println("Temperature in C", temperatureC2)
	// fmt.Println("Temperature in K", temperatureK2)

	// temperatureK3 := 273.15

	// finalTemp := float64(temperatureK3) - 273.15

	// fmt.Println("Final Temperature in C", finalTemp)

	// temperatureC := 16.0
	// var temperatureF float64
	// temperatureF = temperatureC*9/5 + 32

	// fmt.Println(temperatureC, temperatureF)

	// temperatureF := 59.0
	// temperatureC := (temperatureF - 32)*5/9

	// fmt.Println(temperatureC, temperatureF)

	// temperatureK := 288.15
	// temperatureC := temperatureK - 273.15
	// temperatureF := temperatureC*9/5 + 32

	// fmt.Println(temperatureK, temperatureF)

	// temperatureF := 70.0
	// temperatureC := (temperatureF - 32) * 5 / 9
	// temperatureK := temperatureC + 273.15

	// fmt.Println(temperatureF, temperatureK)

	// var temperatureC float64 = 12.15
	// fmt.Println(math.Round(temperatureC))
	// fmt.Println(math.Ceil(temperatureC))  // always round up
	// fmt.Println(math.Floor(temperatureC)) // always round down
	// fmt.Println(math.Abs(-5.5))           // always return positive
	// fmt.Println(math.Pow(3, 2))           // power of firsttosecond
	// fmt.Println(math.Sqrt(16))           // squareroot

	// var tempInt int = 10
	// var tempFloat float64 = float64(tempInt)

	// fmt.Println("interger to float", tempFloat)
	// fmt.Printf("interger to float: %.2f\n", 10.0)

	// fmt.Printf("%T\n", tempInt)
	// fmt.Printf("%T\n", tempFloat)

	// str := fmt.Sprint(80) //convert to string from int because in Go strings are  runes. youcannot string(convert) an int. use fmt.Sprint
	// fmt.Println(str)
	// str2 := strconv.Itoa(10) // or useanotherpackage
	// fmt.Println(str2)

	//STRING TO IN
	// var myStr string = "42"
	// var myIntFromString, _ = strconv.Atoi(myStr) //_ =ignore error value that isreturned
	// fmt.Println("stringto in", myIntFromString)

	//STRING TO FLOAT
	// var floatStr string = "3.14159"
	// var floatFromString, _ = strconv.ParseFloat(floatStr, 64)

	// fmt.Printf("%T\n", floatFromString)

	// myBool, _ := strconv.ParseBool("t")
	// fmt.Println(myBool)

	// const Agency string = "Fast Tracks"

	// const (
	// 	Agency2 = "fast tracks"
	// 	Founded = 2001
	// 	Founder = "James Carter"
	// ) // multiple consts
	// fmt.Println(Agency2)

	// const (
	// 	Economy = iota //iota lets you access index of group of consts
	// 	Compact
	// 	Standard
	// 	FullSize
	// 	Luxury
	// )

	// fmt.Println(Economy)
	// fmt.Println(Compact)
	// fmt.Println(Standard)
	// fmt.Println(FullSize)
	// fmt.Println(Luxury)

	// fmt.Println("whatis yournaem")
	// var name string

	// fmt.Scanln(&name) 
	// // gives umemory address like state & is the pointer to name by reference

	//5000
}
