package main

type Decaf struct {
}

func (b *Decaf) getDescription() string {
	return "Yor order: \n \t Decaf       - 2.49$"
}

func (b *Decaf) cost() float32 {
	return 2.47
}
