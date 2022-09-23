package main

import "fmt"

func main() {
	beverage := &Espresso{}
	espressowithmilk := &Milk{
		Beverage: beverage,
	}
	espressowithmilkandsoy := &Soy{
		Beverage: espressowithmilk,
	}
	fmt.Println(espressowithmilkandsoy.getDescription())
	fmt.Println("Total cost:          -", espressowithmilkandsoy.cost(), "$")

}
