package main

import (
	"fmt"
	"io"
	"os"
)

var points []int = []int{2, 4, 6, 8, 10}

func tryMe(num int) {
	fmt.Printf("You try number is: %v \n", num)
}

func ioTest() {
	// io wrinte string

	const num1, num2, num3 = 5, 10, 15

	// Calling Sprintf() function
	s := fmt.Sprintf("%d + %d = %d", num1, num2, num3)

	// Calling WriteString() function to write the
	// contents of the string "s" to "os.Stdout"
	io.WriteString(os.Stdout, s)

}
