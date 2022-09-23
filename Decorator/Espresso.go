package main

type Espresso struct {
}

func (b *Espresso) getDescription() string {
	return "Yor order: \n \t Espresso    - 1.99$"
}

func (b *Espresso) cost() float32 {
	return 1.99
}
