package main

type Espresso struct {
}

func (b *Espresso) getDescription() string {
	return "Espresso with "
}

func (b *Espresso) cost() float64 {
	return 1.99
}
