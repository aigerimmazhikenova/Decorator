package main

type Mocha struct {
	Beverage Beverage
}

func (b *Mocha) cost() float32 {
	return 0.56 + b.Beverage.cost()
}

func (b *Mocha) getDescription() string {
	return b.Beverage.getDescription() + " \n \t Mocha       - 0.56$"
}
