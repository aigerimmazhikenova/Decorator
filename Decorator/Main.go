package main

import "fmt"

func main() {
	beverage := &DarkRoast{}
	espressowithmilk := &Soy{
		Beverage: beverage,
	}
	espressowithmilkandsoy := &Milk{
		Beverage: espressowithmilk,
	}
	fmt.Println(espressowithmilkandsoy.getDescription())
	cost := fmt.Sprintf("%.2f", espressowithmilkandsoy.cost())
	fmt.Println("Total cost:          -", cost, "$")

}
