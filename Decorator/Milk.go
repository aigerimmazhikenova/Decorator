package main

type Milk struct {
	Beverage Beverage
}

func (b *Milk) cost() float64 {
	return 0.2 + b.Beverage.cost()
}

func (b *Milk) getDescription() string {
	return b.Beverage.getDescription() + "milk, "
}
