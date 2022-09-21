package main

type HouseBlend struct {
}

func (b *HouseBlend) getDescription() string {
	return "HouseBlend with "
}

func (b *HouseBlend) cost() float64 {
	return 1.49
}
