package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	switch option {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item Price: ", reader) // will return a string

		// parse string to float
		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Item added: ", name, p)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount: ", reader)

		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("Tip must be a number")
		}

		b.updateTip(t)
		fmt.Println("Tip updated: ", t)
		fmt.Println("in b ", b.tip)
		promptOptions(b)
	case "s":
		b.saveBill()
		fmt.Println("you save the bill", b.name)
	default:
		fmt.Println("not valid")
		promptOptions(b)
	}
}

func main() {
	var myBill bill = createBill()
	promptOptions(myBill)
}
