package main

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

// var myStringPointer *string // pointer to a string
// var myString string

// myStringPointer = &myString

// func sayHello(s string) {
// 	fmt.Println(s)
// }

// func sayHelloPointer(s *string) {
// 	*s = "Hello, World"
// }

// func main() {
// 	var greeting string = "hello GO"
// 	sayHello(greeting)
// 	fmt.Println("after calling  sayhello,", greeting)
// 	sayHelloPointer(&greeting)
// 	fmt.Println("after calling  sayhelloPointer,", greeting)
// }

// func main () {
// 	var stringPtr *string
// 	var intPtr *int
// 	var floatPtr *float64
// 	var boolPtr *bool

// 	// all return nil
// }

// func main() {
// 	const upperSpeedLimit, lowerSpeedimit int = 120, 80

// 	speed := 120

// 	if speed > upperSpeedLimit {
// 		fmt.Println("too fast")
// 	} else if speed < upperSpeedLimit {
// 		fmt.Println("too slow")
// 	} else {
// 		fmt.Println("hit the gas")
// 	}
// }

// func main() {
// 	color := "red"

// 	switch color  {
// 	case "red":
// 		fmt.Println("red")
// 	case "green":
// 		fmt.Println("green")

// 	default:
// 		fmt.Println("uknown")
// 	}
// }

// func main() {
// 	for i := 0; i < 10; i++ {
// 		if i%2 == 0 {
// 			continue
// 		}
// 		fmt.Println("odd number", i)
// 	}
// }

//while loop
// func main() {
// 	i := 0
// 	for {
// 		fmt.Println("infitnite loop", i)
// 		i++

// 		if i == 3 {
// 			break
// 		}
// 	}
// }

// func main() {
// 	var bodyTypes = [3]string{"sedan", "suv", "convertible"}

// 	// bodyTypes[0] = "sedan"
// 	// bodyTypes[1] = "suv"
// 	// bodyTypes[2] = "convertible"

// 	fmt.Println(bodyTypes)
// }

//2D array
// func main() {
// 	var carFleet [3][2]string
// 	carFleet[0] = [2]string{"5 Sedans available", "2 sedans booked"}
// 	carFleet[1] = [2]string{"3 SUVs available", "4 SUVs booked"}
// 	carFleet[2] = [2]string{"1 convertible available", "1 convertibles booked"}

// 	fmt.Println("carfleet status:")

// 	for i := 0; i < len(carFleet); i++ {
// 		row := carFleet[i]
// 		for j := 0; j < len(row); j++ {
// 			fmt.Printf("%v ", row[j])
// 		}
// 		fmt.Println()
// 	}
// }

//SLICES
// func main() {
// 	var fuelTypes = []string{"Electric", "Gasoline", "Hybrid"}
// 	fmt.Println(fuelTypes)

// 	fuelTypes = append(fuelTypes, "Diesel", "hydrogen")
// 	fmt.Println(fuelTypes)
// }

// func main() {
// 	fuelTypes := make([]string, 3) // makesarray with predeifne len 3
// 	fuelTypes = append(fuelTypes, "one", "two", "three")
// 	fuelTypes[0] = "electric"
// 	fuelTypes[1] = "diesel"
// 	fmt.Println(fuelTypes)

// 	fuelTypesCopy := make([]string, len(fuelTypes))
// 	copy(fuelTypesCopy, fuelTypes)
// 	fmt.Println(fuelTypesCopy)
// }

// func main() {
// 	fuelTypes := []string{"gas", "diesel", "electric", "hybrid", "hydrogen"}

// 	popular := fuelTypes[0:2]

// 	fmt.Println("popular:", popular)

// 	clean := fuelTypes[2:]
// 	fmt.Println("clean:", clean)

// 	electric := fuelTypes[2:4]
// 	fmt.Println("electric:", electric)
// }

//MAPS
// func main() {
// 	carInventory := map[string]int{

// 		"Sedan":       25,
// 		"SUV":         15,
// 		"Convertible": 10,
// 	}

// fmt.Println("car inventory:", carInventory)
// fmt.Println("types", len(carInventory))

// numberofSedans := carInventory["Sedan"]
// fmt.Printf("we have %v sedans\n", numberofSedans)

// carInventory["Sedan"] = 20

// numberofSedans2 := carInventory["Sedan"]
// fmt.Printf("we have %v sedans\n", numberofSedans2)

// delete(carInventory, "Convertible")

// fmt.Println(carInventory)

// numSedans, sedanFound := carInventory["Sedan"]
// fmt.Println("sedans found", sedanFound)

// if sedanFound {
// 	fmt.Println("We have", numSedans)
// }

//cCOMMON PATTERN FOR MAPS
// if numSedans, ok := carInventory["Sedan"]; ok { // ok is just avaribale name to check if carinventory sedan istruthy. the second ok is if true
// 	fmt.Println("we have", numSedans, "Sedans")
// }

//FOR RANGE LOOP FOR ARRAYS SLICES AND MAPS

// func main() {
// 	bodyTypes := [3]string{"Sedan", "SUV", "Convertible"}

// 	for i, bodyType := range bodyTypes {
// 		fmt.Printf("Index: %v. Item: %v\n", i, bodyType)
// 	}

// 	//this captures indexes only
// 	for i := range bodyTypes {
// 		fmt.Printf("Index: %v\n", i)
// 	}

// 	//forvalue only notindex

// 	for _, bodyType:= range bodyTypes {
// 		fmt.Printf("Index: %v\n", bodyType)
// 	}

// }

//RANGE LOOP FOR MAPS
// func main() {

// 	carInventory := map[string]int{
// 		"Sedan":       25,
// 		"SUV":         7,
// 		"Convertible": 10,
// 		"Hatchback":   8,
// 	}

// 	//to call both value and key
// 	for bodyType, count := range carInventory {
// 		fmt.Printf("%v -> %v\n", bodyType, count)
// 	}

// 	//KEY ONLY
// 	for bodyType := range carInventory {
// 		fmt.Printf("%v\n", bodyType)
// 	}
// 	//VALUE ONLY
// 	totalInventory := 0
// 	for _, count := range carInventory {
// 		totalInventory += count
// 	}
// 	fmt.Printf("%v cars in total\n", totalInventory)
// }

//FUNCTIONS

// var carInventory = map[string]int{}

// func addToInventory(bodyType string, count int) {
// 	carInventory[bodyType] += count
// 	fmt.Println(bodyType, "added. New count:", carInventory[bodyType])
// }
// func main() {
// 	addToInventory("Sedan", 10)
// 	addToInventory("SUV", 7)
// 	addToInventory("Sedan", 2)
// 	fmt.Println("Updated Car Inventory", carInventory)
// }

//RETURNING VARIABLES
// var carInventory = map[string]int{}

// func addToInventory(bodyType string, count int) {
// 	carInventory[bodyType] += count
// 	fmt.Println(bodyType, "added. New count:", carInventory[bodyType])
// }

// func getCount(bodyType string) int {
// 	fmt.Printf("looking up %v", bodyType)
// 	count := carInventory[bodyType]
// 	return count
// }

// func main() {
// 	addToInventory("Coupe", 3)
// 	fmt.Println("Found:", getCount("Coupe"))
// }

//RETURNING MULTIPLE TYPES OF VARIABLES

// var carInventory = map[string]int{}

// func getCount(bodyType string) int {
// 	count := carInventory[bodyType]
// 	return count
// }

// func getStatus(bodyType string) (available string, booked string) {
// 	status := carFleet[bodyType]
// 	return available, booked
// }

// func main() {
// 	addToFleet("Sedan", 5, 2)
// 	available, booked := getStatus("Sedan")
// 	fmt.Println(available)
// 	fmt.Println(booked)
// }

//VARIADIC FUNCTION - accepts ininifite arguments
// import "fmt"

// func sum(numbers ...int) int {
// 	total := 0
// 	for _, n := range numbers {
// 		total += n
// 	}

// 	return total
// }

// func main() {
// 	fmt.Println(sum(1, 2, 3, 43, 4, 4))
// }

//CLOSURES

// func sequence() func() int {

// 	i := 0
// 	return func() int {
// 		i++
// 		return i
// 	}
// }

// func main() {
// 	nextInt := sequence()
// 	fmt.Println(nextInt())
// 	fmt.Println(nextInt())
// 	fmt.Println(nextInt())

// 	newInts := sequence()
// 	fmt.Println(newInts())
// }

//FIBONACCI SEQUENCE

// func fib() func() int {
// 	a, b := 0, 1
// 	return func() int {
// 		a, b = b, a+b // THIS IS SAYING (A, B) = (ASSIGN A=B, B=A+B)
// 		return a
// 	}
// }

// func main() {
// 	f := fib()
// 	fmt.Println(f(), f(), f(), f(), f())
// }

// //STRUCTS
// import (
// 	"errors"
// 	"fmt"
// 	"time"
// )

// type AuditInfo struct {
// 	CreatedAt    time.Time
// 	LastModified time.Time
// }

// type BankAccount struct {
// 	AccountNumber AccountNumber
// 	Balance       float64
// 	AuditInfo
// }

// type Customer struct {
// 	Name     string
// 	accounts []*BankAccount
// 	AuditInfo
// }

// type Bank struct {
// 	Name      string
// 	Customers []*Customer
// }

// type AccountNumber string
// type Balance = float64 //type alias assigns a name toa float so ucan write func (balance balance) instead of (balance float64)

// // putting the name aftert he argument makes theargument a reciever. binds it to the fucntion displayBalance so  you can call it as account.displayBalnace() instaed ofdisplayBalnace(account)
// func (ba BankAccount) DisplayBalance() {
// 	fmt.Printf("Account number: %s, Balance: $%.2f\n", ba.AccountNumber, ba.Balance)
// }

// func (ba *BankAccount) Deposit(amount float64) {
// 	ba.Balance += amount
// }

// func (ba *BankAccount) Withdraw(amount float64) error {
// 	if ba.Balance < amount {
// 		return errors.New("NSF")
// 	}
// 	ba.Balance -= amount
// 	return nil
// }

// func (c *Customer) AddAccount(account *BankAccount) {
// 	c.accounts = append(c.accounts, account)
// }

// func (c Customer) DisplayAccounts() {
// 	fmt.Printf("Customer: %s\n", c.Name)
// 	for _, account := range c.accounts {
// 		account.DisplayBalance()
// 	}
// }

// func NewBank(name string) *Bank {
// 	return &Bank{Name: "GoBank"}
// }

// func NewCustomer(name string) *Customer {
// 	return &Customer{
// 		Name:      name,
// 		AuditInfo: AuditInfo{CreatedAt: time.Now(), LastModified: time.Now()},
// 	}
// }

// func NewBankAccount(num AccountNumber) *BankAccount {
// 	return &BankAccount{
// 		AccountNumber: num,
// 		AuditInfo:     AuditInfo{CreatedAt: time.Now(), LastModified: time.Now()},
// 	}
// }

// func (a AccountNumber) isValid() bool {
// 	return len(string(a)) == 10
// }

// func main() {

// bank := NewBank("Go Bank")
// customer1 := NewCustomer("Alice")
// customer2 := NewCustomer("Bob")

// account1 := NewBankAccount("834354")
// account1.Deposit(1000.00)

// account2 := NewBankAccount("8343345")
// account2.Deposit(5000.00)

// customer1.AddAccount(*account1)
// customer2.AddAccount(*account2)

// bank.AddCustomer(*customer1)
// bank.AddCustomer(*customer2)

// var acct AccountNumber
// fmt.Printf("acct is %T\n", acct)

// var acct2 string
// fmt.Printf("acct2 is %T\n", acct)

// account := NewBankAccount("4234234")
// fmt.Println("THIS IS", account.AccountNumber.isValid())
// 	account := BankAccount{AccountNumber: "12324324", Balance: 1000.00}
// 	fmt.Println(account.Balance, account.AccountNumber)

// 	account.DisplayBalance()

// 	// account2 := &BankAccount{AccountNumber: "8978945", Balance: 500.99}
// 	// account3 := &BankAccount{AccountNumber: "89234234234235", Balance: 6500.99, AuditInfo: AuditInfo{createdAt: time.Now(), LastModified: time.Now()}}
// 	account.DisplayBalance()
// 	account.Deposit(1000)
// 	account.DisplayBalance()

// 	//add accts to customer
// 	customer1.AddAccount(account1)
// 	customer2.AddAccount(account2)

// 	//add customers tothebank

// bank.AddCustomer(customer1)
// bank.AddCustomer(customer2)

// bank.DisplayCustomers()
// fmt.Println(account3.createdAt)
// }

//238 interfaces just let data flow through unlike structswhich store data
import "fmt"

// 1️⃣ Define an interface
type Greeter interface {
	Greet()
}

// 2️⃣ Person type
type Person struct {
	Name string
}

// 3️⃣ Robot type
type Robot struct {
	ID int
}

//IN ORDER TO MAKE INTERFACE WORK AS A REUSABLE FUNCITON WITH MULTIPLE TYPES, YOU NEED TO DECLARE  TEH SAME METHOD WITHIN FOR ALL POSSIBLE CASES USING THE METHOD NAME
func (p Person) Greet() {
	fmt.Println("Hi, I'm", p.Name)
}

// Robot implements Greet()
func (r Robot) Greet() {
	fmt.Println("Beep boop, I am robot #", r.ID)
}

// 4️⃣ Function that works with any Greeter
func SayHello(g Greeter) {
	g.Greet() // just call the method
}

func main() {
	alice := Person{"Alice"}
	r2d2 := Robot{42}

	SayHello(alice) // Hi, I'm Alice
	SayHello(r2d2)  // Beep boop, I am robot # 42
}
