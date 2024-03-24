package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, error := r.ReadString('\n')

	return strings.TrimSpace(input), error
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	bill := newBill(name)

	fmt.Println("Created the bill - ", bill.name)

	return bill
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	option, _ := getInput("Choose option (a - add, s - save, t - add tip) ", reader)

	fmt.Println(option)
}

func main() {
	var myBill bill = createBill()
	promptOptions(myBill)
	fmt.Println("The bill", myBill)
}
