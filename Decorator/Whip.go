package main

type Whip struct {
	Beverage Beverage
}

func (b *Whip) cost() float32 {
	return 0.39 + b.Beverage.cost()
}

func (b *Whip) getDescription() string {
	return b.Beverage.getDescription() + " \n \t Whip        - 0.39$"
}
