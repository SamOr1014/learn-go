package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func twoArg(n []string, f func(string)) {
	for _, val := range n {
		f(val)
	}
}

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func returnTwoVal(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func updateMap(mapping map[string]float64) {
	mapping["cls"] = 0.99
}

func updateString(string string) {
	string = "not tifa"
}

func updateStringViaPointer(ptr *string) {
	*ptr = "not tifa"
}

func main() {
	fmt.Println("\n____________Declaring Variable______________\n")
	// string
	var nameOne string = "Sam"

	var nameTwo = "Elise"

	var combine string = nameOne + nameTwo

	fmt.Println(combine)

	nameOne = "Haw"
	nameThree := "Yee"

	combine = nameOne + nameThree

	fmt.Println(combine)

	// int
	integer := 26

	var integerTwo int = 25

	var add = integer + integerTwo

	fmt.Println(add)
	//  bit and memory int
	var numWithBit int8 = 127
	var numWithBitUnsigned uint8 = 255
	var numWithBitNeg int8 = -128
	var numWithBit16 int16 = 32767

	fmt.Println("\n____________Print and format______________\n")
	// fmt printing
	fmt.Println(numWithBit, numWithBitUnsigned, numWithBitNeg, numWithBit16)
	fmt.Printf("Float is %0.3f \n", 2.75421)

	var savedString string = fmt.Sprintf("%v and %v is good", "taco", "rice")
	fmt.Println(savedString)

	fmt.Println("\n____________Array______________\n")
	// array
	var nums [4]int = [4]int{1, 2, 3, 4}
	// another array declare
	var numDeclare = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	stringDeclare := [3]string{"sam", "elise", "nothing"}
	stringDeclare[2] = "sth"
	fmt.Printf("nums length: %v %v \n", len(nums), len(numDeclare))
	fmt.Println("string array len and content: ", len(stringDeclare), stringDeclare)

	fmt.Println("\n____________Slices______________\n")
	// slices
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice[4] = slice[7] + 2
	fmt.Println("slice", slice)
	fmt.Println("appended slice", append(slice, 9))
	// slice range
	sliceRange := slice[2:6]
	fmt.Println("slice range", sliceRange)
	fmt.Println("slice from 1", slice[1:])

	fmt.Println("\n_____________Std lib______________\n")
	// Standard Library

	ioTest()
	sampleString := "Yeeee haw"
	// These function doesn't mutate the original string
	fmt.Println(strings.Contains(sampleString, "haw"))            //true
	fmt.Println(strings.ReplaceAll(sampleString, "Yeeee", "Yee")) //Yee Haw
	fmt.Println(strings.ToUpper(sampleString))                    // YEEEE HAW
	fmt.Println(strings.Index(sampleString, "aw"))                // 7
	fmt.Println(strings.Split(sampleString, " "))                 // [Yeee haw]

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	fmt.Println("unsorted ages", ages)
	// This sort will mutate the original slice
	sort.Ints(ages)
	fmt.Println("sorted ages", ages)

	index := sort.SearchInts(ages, 30)
	index2 := sort.SearchInts(ages, 100)
	fmt.Println("index", index) // 2
	// if number not exist in the slice will return length of slice
	fmt.Println("index", index2, "length of ages", len(ages)) // 8 === len(ages)

	name := []string{"a", "y", "b", "u", "c"}
	sort.Strings(name)
	fmt.Println("sorted", name)

	fmt.Println(sort.SearchStrings(name, "c")) // 2

	fmt.Println("\n_____________Loops______________\n")
	x := 0
	// while loop like
	for x < 5 {
		fmt.Println("val of x: ", x)
		x++
	}
	// normal for loop
	for i := 0; i < 5; i++ {
		fmt.Println("val of i: ", i)
	}

	// looping slice
	for i := 0; i < len(name); i++ {
		fmt.Println("name", name[i])
	}
	fmt.Println("for in loops of", name)
	// for in loop
	for index, value := range name {
		fmt.Printf("index of cur: %v; value of cur: %v \n", index, value)
	}
	// value only example vice versa for index
	for _, value := range name {
		fmt.Printf("value of cur: %v \n", value)
	}
	//update value example plus a condition example
	for idx, _ := range name {
		if idx == len(name)-1 {
			name[idx] = "lastOne"
		} else if idx == 0 {
			name[idx] = "first"
		} else {
			fmt.Println("doesn't hit")
		}
	}
	fmt.Println("re assign val", name) // [first b c u lastOne]
	for idx, val := range name {
		if idx == 1 {
			fmt.Println("index is: ", idx, "and shd continue")
			continue
		}
		if val == "c" {
			fmt.Println("break at c")
			break
		}
		fmt.Println("pass all if and do nothing :D")
	}

	fmt.Println("\n_____________declare funcs______________\n")
	sayGreeting("Sam") //Good morning Sam
	sayBye("Elise")    //Goodbye Elise
	twoArg([]string{"Sam", "Elise"}, sayGreeting)
	twoArg([]string{"Sam", "Elise"}, sayBye)
	c1 := circleArea(10.5)
	c2 := circleArea(15)
	fmt.Println("circle", c1, c2)
	fmt.Printf("c1 is %0.3f and c2 is %0.3f \n", c1, c2)

	// return two val
	firstN, secondN := returnTwoVal("Sam Or")
	fmt.Println("First Name: ", firstN)
	fmt.Println("Last Name: ", secondN)

	// package scope -> have to run `go run *.go`/`go run main.go package.go`
	tryMe(1)
	for _, val := range points {
		tryMe(val)
	}

	fmt.Println("\n_____________map______________\n")
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee-pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["soup"])
	// looping map
	for key, val := range menu {
		fmt.Println(key, "-", val)
	}
	// ints as key type
	phoneBook := map[int]string{
		12345678: "sam",
		24681012: "elise",
		36911131: "mom",
	}
	fmt.Println("phoneBook", phoneBook)
	fmt.Println("phoneBook", phoneBook[12345678])
	// update map val
	phoneBook[12345678] = "sam2"
	fmt.Println("phoneBook", phoneBook)

	// pass by value feat
	// Go is pass by value lang , int string float struct bools arrays will be copy after passing into a function
	// mean while map slice and functions are not thy are passed by ref (to their addr)
	// that means u can change val up for map and slice in a function
	wannaUpdate := "tifa"
	updateString(wannaUpdate)
	fmt.Println(wannaUpdate) // supposedly shd be still tifa
	updateMap(menu)
	fmt.Println("updated menu", menu) // cls shd be in and priced 0.99

	// getting the pointer to the memory addr o the var
	storePointer := &wannaUpdate
	fmt.Println("Address is: ", storePointer)
	fmt.Println("de-ref a pointer and get a val: ", *storePointer)

	// here we pass in the pointer to wannaUpdate and update the val inside
	fmt.Println("before passing in the pointer: ", wannaUpdate)
	updateStringViaPointer(&wannaUpdate)
	fmt.Println("Val will update after pointer passed in: ", wannaUpdate) // expect: not tifa

	// rule of thumb : val -> addr : &____ ; addr -> val : *____

	fmt.Println("\n_____________struct______________\n")

	// checkout bill.go for implementation
	myBill := newBill("sam's bill")
	fmt.Println("my bill: ", myBill)
	fmt.Println("my bill name: ", myBill.name)
	fmt.Println("my bill items: ", myBill.item)
	fmt.Println("my bill tip: ", myBill.tip)

	fmt.Println("\n_____________Receiver func______________\n")
	// checkout bill.go for implementation

	// print out the formatted strin
	myBill.updateTip(10.2)
	myBill.addItem("kebab", 5.50)
	myBill.addItem("tea", 0.89)
	fmt.Println(myBill.formatBill())

}
