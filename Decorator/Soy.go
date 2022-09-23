package main

type Soy struct {
	Beverage Beverage
}

func (b *Soy) cost() float32 {
	return 0.35 + b.Beverage.cost()
}

func (b *Soy) getDescription() string {
	return b.Beverage.getDescription() + "\n \t Soy         - 0.35$ "
}
