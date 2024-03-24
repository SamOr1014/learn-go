package main

import "fmt"

type bill struct {
	name string
	item map[string]float64
	tip  float64
}

func newBill(name string) bill {
	b := bill{
		name: name,
		item: map[string]float64{"pie": 1.99, "cake": 2.99},
		tip:  0,
	}
	return b
}

func (b *bill) formatBill() string {
	formattedString := "Bill breakdown: \n"
	var total float64 = 0

	for key, val := range b.item {
		formattedString += fmt.Sprintf("%-25v ...$%v \n", key+":", val)
		total += val
	}

	formattedString += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	formattedString += fmt.Sprintf("%-25v ...$%0.2f\n", "total:", total+b.tip)

	return formattedString
}

func (b *bill) updateTip(tip float64) {
	(*b).tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.item[name] = price
}
