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

	fmt.Println("Yor order:")
	fmt.Println(espressowithmilkandsoy.getDescription())
	fmt.Println("Cost: ", espressowithmilkandsoy.cost(), "$")

}
