package main

type Whip struct {
	Beverage Beverage
}

func (b *Whip) cost() float64 {
	return 0.39 + b.Beverage.cost()
}

func (b *Whip) getDescription() string {
	return b.Beverage.getDescription() + "whip, "
}
