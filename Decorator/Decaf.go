package main

type Decaf struct {
}

func (b *Decaf) getDescription() string {
	return "Decaf with "
}

func (b *Decaf) cost() float64 {
	return 2.49
}
