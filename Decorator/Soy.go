package main

type Soy struct {
	Beverage Beverage
}

func (b *Soy) cost() float64 {
	return 0.35 + b.Beverage.cost()
}

func (b *Soy) getDescription() string {
	return b.Beverage.getDescription() + "soy, "
}
