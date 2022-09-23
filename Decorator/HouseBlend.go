package main

type HouseBlend struct {
}

func (b *HouseBlend) getDescription() string {
	return "Yor order: \n \t HouseBlend  - 1.49$"
}

func (b *HouseBlend) cost() float32 {
	return 1.49
}
