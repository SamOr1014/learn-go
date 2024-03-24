package main

import "fmt"

type bill struct {
	name string
	item map[string]float64
	tip  float64
}

// make new bill
func newBill(name string) bill {
	b := bill{
		name: name,
		item: map[string]float64{"pie": 1.99, "cake": 2.99},
		tip:  0,
	}
	return b
}

// format bill
// Another reason to pass in pointer instead of the val is to
// prevent making copy of the original data
// more memory efficient
func (b *bill) formatBill() string {
	formattedString := "Bill breakdown: \n"
	var total float64 = 0

	for key, val := range b.item {
		formattedString += fmt.Sprintf("%-25v ...$%v \n", key+":", val)
		total += val
	}
	// add tip
	formattedString += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)
	// total
	formattedString += fmt.Sprintf("%-25v ...$%0.2f\n", "total:", total+b.tip)

	return formattedString
}

// update tip
func (b *bill) updateTip(tip float64) {
	// can de ref (*b) or it will auto de ref the pointer in receiver func
	// can see addItem
	(*b).tip = tip
}

// add item to the bill
func (b *bill) addItem(name string, price float64) {
	b.item[name] = price
}
